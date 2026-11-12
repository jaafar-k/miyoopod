package main

import "testing"

func TestShapeForDisplayArabicJoins(t *testing.T) {
	// "باب" should use beginning/medial/final presentation forms, then reverse for LTR draw.
	got := shapeForDisplay("باب")
	if got == "باب" {
		t.Fatalf("expected shaped output, got logical %q", got)
	}
	if !containsArabic(got) {
		t.Fatalf("expected Arabic presentation forms, got %q", got)
	}
	// Idempotent: shaping again must not change visual order.
	if again := shapeForDisplay(got); again != got {
		t.Fatalf("shaping not idempotent: %q -> %q", got, again)
	}
}

func TestShapeForDisplayLatinUnchanged(t *testing.T) {
	in := "Hello World"
	if got := shapeForDisplay(in); got != in {
		t.Fatalf("latin changed: %q", got)
	}
}

func TestShapeMixedKeepsLatinOrder(t *testing.T) {
	in := "Artist - سلام"
	got := shapeForDisplay(in)
	if got[:6] != "Artist" {
		t.Fatalf("expected Latin prefix preserved, got %q", got)
	}
}

func TestLamAlefLigature(t *testing.T) {
	got := shapeForDisplay("لا")
	runes := []rune(got)
	if len(runes) != 1 {
		t.Fatalf("expected lam-alef ligature as one glyph, got %q (%d runes)", got, len(runes))
	}
}
