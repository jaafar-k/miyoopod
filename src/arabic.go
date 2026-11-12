package main

import (
	"unicode"

	"golang.org/x/image/font"
)

// shapeForDisplay converts logical Arabic text into a left-to-right visual
// string with joined presentation forms so FreeType/gg can render it.
// Idempotent: already-shaped presentation-form strings are returned as-is.
func shapeForDisplay(s string) string {
	if s == "" || !needsArabicShaping(s) {
		return s
	}
	rtl := isRTLDominant(s)
	shaped := arabicToGlyphs(s)
	return arabicBidiVisual(shaped, rtl)
}

// needsArabicShaping is true when s still has logical Arabic letters to join.
func needsArabicShaping(s string) bool {
	for _, r := range s {
		if isLogicalArabicLetter(r) {
			return true
		}
	}
	return false
}

// containsArabic is true for logical Arabic or presentation-form glyphs.
func containsArabic(s string) bool {
	for _, r := range s {
		if isArabicRune(r) {
			return true
		}
	}
	return false
}

func isLogicalArabicLetter(r rune) bool {
	switch {
	case r >= 0x0600 && r <= 0x06FF:
		return true
	case r >= 0x0750 && r <= 0x077F:
		return true
	case r >= 0x08A0 && r <= 0x08FF:
		return true
	default:
		return false
	}
}

func isArabicRune(r rune) bool {
	if isLogicalArabicLetter(r) {
		return true
	}
	switch {
	case r >= 0xFB50 && r <= 0xFDFF: // Presentation Forms-A
		return true
	case r >= 0xFE70 && r <= 0xFEFF: // Presentation Forms-B
		return true
	default:
		return false
	}
}

func isRTLDominant(s string) bool {
	for _, r := range s {
		if isLogicalArabicLetter(r) || (r >= 0xFB50 && r <= 0xFDFF) || (r >= 0xFE70 && r <= 0xFEFF) {
			return true
		}
		if unicode.Is(unicode.Latin, r) || (r >= '0' && r <= '9') {
			return false
		}
	}
	return false
}

func arabicToGlyphs(s string) string {
	runes := applyLamAlefLigatures([]rune(s))
	out := make([]rune, 0, len(runes))
	for i, cur := range runes {
		if arabicTashkeel[cur] {
			out = append(out, cur)
			continue
		}
		prev := prevJoiningRune(runes, i)
		next := nextJoiningRune(runes, i)
		out = append(out, arabicGlyphForm(prev, cur, next))
	}
	return string(out)
}

func applyLamAlefLigatures(in []rune) []rune {
	out := make([]rune, 0, len(in))
	for i := 0; i < len(in); i++ {
		cur := in[i]
		if cur != arLam.Unicode && cur != arLam.Beginning && cur != arLam.Medium &&
			cur != arLam.Final && cur != arLam.Isolated {
			out = append(out, cur)
			continue
		}
		j := i + 1
		for j < len(in) && arabicTashkeel[in[j]] {
			j++
		}
		if j >= len(in) {
			out = append(out, cur)
			continue
		}
		lig, ok := lamAlefLigature(in[j])
		if !ok {
			out = append(out, cur)
			continue
		}
		out = append(out, lig)
		// Preserve any tashkeel that sat between lam and alef.
		out = append(out, in[i+1:j]...)
		i = j
	}
	return out
}

func lamAlefLigature(alef rune) (rune, bool) {
	switch alef {
	case arAlef.Unicode, arAlef.Isolated, arAlef.Final, arAlef.Medium:
		return arLamAlef.Unicode, true
	case arAlefHamzaAbove.Unicode, arAlefHamzaAbove.Isolated, arAlefHamzaAbove.Final, arAlefHamzaAbove.Medium:
		return arLamAlefHamzaA.Unicode, true
	case arAlefHamzaBelow.Unicode, arAlefHamzaBelow.Isolated, arAlefHamzaBelow.Final, arAlefHamzaBelow.Medium:
		return arLamAlefHamzaB.Unicode, true
	case arAlefMadda.Unicode, arAlefMadda.Isolated, arAlefMadda.Final, arAlefMadda.Medium:
		return arLamAlefMadda.Unicode, true
	default:
		return 0, false
	}
}

func prevJoiningRune(runes []rune, i int) rune {
	for j := i - 1; j >= 0; j-- {
		if arabicTashkeel[runes[j]] {
			continue
		}
		return runes[j]
	}
	return 0
}

func nextJoiningRune(runes []rune, i int) rune {
	for j := i + 1; j < len(runes); j++ {
		if arabicTashkeel[runes[j]] {
			continue
		}
		return runes[j]
	}
	return 0
}

func arabicGlyphForm(prev, cur, next rune) rune {
	h, ok := lookupArabicHarf(cur)
	if !ok {
		return cur
	}
	prevOK := false
	nextOK := false
	if _, ok := lookupArabicHarf(prev); ok {
		prevOK = true
	}
	if _, ok := lookupArabicHarf(next); ok {
		nextOK = true
	}

	connectFromPrev := prevOK && !arabicRightJoinOnly[canonicalArabic(prev)]
	if prevOK && nextOK {
		if arabicRightJoinOnly[canonicalArabic(prev)] {
			return h.Beginning
		}
		return h.Medium
	}
	if nextOK {
		return h.Beginning
	}
	if connectFromPrev {
		return h.Final
	}
	return h.Isolated
}

func canonicalArabic(r rune) rune {
	if h, ok := lookupArabicHarf(r); ok {
		return h.Unicode
	}
	return r
}

// arabicBidiVisual reorders shaped text for LTR drawing.
// RTL-dominant strings are fully reversed; mixed strings reverse Arabic runs only.
func arabicBidiVisual(s string, rtlDominant bool) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return s
	}
	if rtlDominant {
		reverseRunes(runes)
		mirrorRTLRunes(runes)
		return string(runes)
	}
	return string(reverseArabicRuns(runes))
}

func reverseArabicRuns(runes []rune) []rune {
	out := make([]rune, 0, len(runes))
	for i := 0; i < len(runes); {
		if !isArabicRune(runes[i]) && !arabicTashkeel[runes[i]] {
			out = append(out, runes[i])
			i++
			continue
		}
		j := i
		for j < len(runes) {
			r := runes[j]
			if isArabicRune(r) || arabicTashkeel[r] {
				j++
				continue
			}
			// Keep neutral chars inside an Arabic run when Arabic continues after.
			if isArabicNeutral(r) && arabicContinues(runes, j+1) {
				j++
				continue
			}
			break
		}
		seg := append([]rune(nil), runes[i:j]...)
		reverseRunes(seg)
		mirrorRTLRunes(seg)
		out = append(out, seg...)
		i = j
	}
	return out
}

func arabicContinues(runes []rune, i int) bool {
	for ; i < len(runes); i++ {
		if arabicTashkeel[runes[i]] || isArabicNeutral(runes[i]) {
			continue
		}
		return isArabicRune(runes[i])
	}
	return false
}

func isArabicNeutral(r rune) bool {
	switch r {
	case ' ', '\t', '-', '–', '—', '_', '.', ',', ':', ';', '!', '?',
		'/', '\\', '|', '"', '\'', '«', '»', '(', ')', '[', ']', '{', '}',
		'،', '؛', '؟', 'ـ':
		return true
	default:
		return false
	}
}

func reverseRunes(runes []rune) {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

var arabicMirrorPairs = map[rune]rune{
	'(': ')', ')': '(',
	'[': ']', ']': '[',
	'{': '}', '}': '{',
	'<': '>', '>': '<',
	'«': '»', '»': '«',
}

func mirrorRTLRunes(runes []rune) {
	for i, r := range runes {
		if m, ok := arabicMirrorPairs[r]; ok {
			runes[i] = m
		}
	}
}

// fontForText returns the Arabic-capable face when text needs Arabic glyphs.
func (app *MiyooPod) fontForText(text string, face font.Face) font.Face {
	if !containsArabic(text) || !app.ArabicFontReady {
		return face
	}
	switch face {
	case app.FontHeader:
		return app.FontHeaderAR
	case app.FontMenu:
		return app.FontMenuAR
	case app.FontTitle:
		return app.FontTitleAR
	case app.FontArtist:
		return app.FontArtistAR
	case app.FontAlbum:
		return app.FontAlbumAR
	case app.FontTime:
		return app.FontTimeAR
	case app.FontSmall:
		return app.FontSmallAR
	case app.FontHeaderAR, app.FontMenuAR, app.FontTitleAR, app.FontArtistAR,
		app.FontAlbumAR, app.FontTimeAR, app.FontSmallAR:
		return face
	default:
		return app.FontMenuAR
	}
}

