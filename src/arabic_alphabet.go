package main

// Arabic letter forms for presentation-form shaping (FreeType/gg have no HarfBuzz).
// Letter tables adapted from github.com/01walid/goarabic (MIT).

type arabicHarf struct {
	Unicode, Isolated, Beginning, Medium, Final rune
}

var (
	arFatha    = '\u064e'
	arFathatan = '\u064b'
	arDamma    = '\u064f'
	arDammatan = '\u064c'
	arKasra    = '\u0650'
	arKasratan = '\u064d'
	arShadda   = '\u0651'
	arSukun    = '\u0652'
)

var arabicTashkeel = map[rune]bool{
	arFatha: true, arFathatan: true, arDamma: true, arDammatan: true,
	arKasra: true, arKasratan: true, arShadda: true, arSukun: true,
}

var (
	arAlefHamzaAbove = arabicHarf{'\u0623', '\ufe83', '\u0623', '\ufe84', '\ufe84'}
	arAlef           = arabicHarf{'\u0627', '\ufe8d', '\u0627', '\ufe8e', '\ufe8e'}
	arAlefMadda      = arabicHarf{'\u0622', '\ufe81', '\u0622', '\ufe82', '\ufe82'}
	arHamza          = arabicHarf{'\u0621', '\ufe80', '\u0621', '\u0621', '\u0621'}
	arWawHamza       = arabicHarf{'\u0624', '\ufe85', '\u0624', '\ufe86', '\ufe86'}
	arAlefHamzaBelow = arabicHarf{'\u0625', '\ufe87', '\u0625', '\ufe88', '\ufe88'}
	arYehHamza       = arabicHarf{'\u0626', '\ufe89', '\ufe8b', '\ufe8c', '\ufe8a'}
	arBeh            = arabicHarf{'\u0628', '\ufe8f', '\ufe91', '\ufe92', '\ufe90'}
	arPeh            = arabicHarf{'\u067e', '\ufb56', '\ufb58', '\ufb59', '\ufb57'}
	arTeh            = arabicHarf{'\u062a', '\ufe95', '\ufe97', '\ufe98', '\ufe96'}
	arTehMarbuta     = arabicHarf{'\u0629', '\ufe93', '\u0629', '\u0629', '\ufe94'}
	arTheh           = arabicHarf{'\u062b', '\ufe99', '\ufe9b', '\ufe9c', '\ufe9a'}
	arJeem           = arabicHarf{'\u062c', '\ufe9d', '\ufe9f', '\ufea0', '\ufe9e'}
	arTcheh          = arabicHarf{'\u0686', '\ufb7a', '\ufb7c', '\ufb7d', '\ufb7b'}
	arHah            = arabicHarf{'\u062d', '\ufea1', '\ufea3', '\ufea4', '\ufea2'}
	arKhah           = arabicHarf{'\u062e', '\ufea5', '\ufea7', '\ufea8', '\ufea6'}
	arDal            = arabicHarf{'\u062f', '\ufea9', '\u062f', '\ufeaa', '\ufeaa'}
	arThal           = arabicHarf{'\u0630', '\ufeab', '\u0630', '\ufeac', '\ufeac'}
	arReh            = arabicHarf{'\u0631', '\ufead', '\u0631', '\ufeae', '\ufeae'}
	arJeh            = arabicHarf{'\u0698', '\ufb8a', '\u0698', '\ufb8b', '\ufb8b'}
	arZain           = arabicHarf{'\u0632', '\ufeaf', '\u0632', '\ufeb0', '\ufeb0'}
	arSeen           = arabicHarf{'\u0633', '\ufeb1', '\ufeb3', '\ufeb4', '\ufeb2'}
	arSheen          = arabicHarf{'\u0634', '\ufeb5', '\ufeb7', '\ufeb8', '\ufeb6'}
	arSad            = arabicHarf{'\u0635', '\ufeb9', '\ufebb', '\ufebc', '\ufeba'}
	arDad            = arabicHarf{'\u0636', '\ufebd', '\ufebf', '\ufec0', '\ufebe'}
	arTah            = arabicHarf{'\u0637', '\ufec1', '\ufec3', '\ufec4', '\ufec2'}
	arZah            = arabicHarf{'\u0638', '\ufec5', '\ufec7', '\ufec8', '\ufec6'}
	arAin            = arabicHarf{'\u0639', '\ufec9', '\ufecb', '\ufecc', '\ufeca'}
	arGhain          = arabicHarf{'\u063a', '\ufecd', '\ufecf', '\ufed0', '\ufece'}
	arFeh            = arabicHarf{'\u0641', '\ufed1', '\ufed3', '\ufed4', '\ufed2'}
	arQaf            = arabicHarf{'\u0642', '\ufed5', '\ufed7', '\ufed8', '\ufed6'}
	arKaf            = arabicHarf{'\u0643', '\ufed9', '\ufedb', '\ufedc', '\ufeda'}
	arKeheh          = arabicHarf{'\u06a9', '\ufb8e', '\ufb90', '\ufb91', '\ufb8f'}
	arGaf            = arabicHarf{'\u06af', '\ufb92', '\ufb94', '\ufb95', '\ufb93'}
	arLam            = arabicHarf{'\u0644', '\ufedd', '\ufedf', '\ufee0', '\ufede'}
	arMeem           = arabicHarf{'\u0645', '\ufee1', '\ufee3', '\ufee4', '\ufee2'}
	arNoon           = arabicHarf{'\u0646', '\ufee5', '\ufee7', '\ufee8', '\ufee6'}
	arHeh            = arabicHarf{'\u0647', '\ufee9', '\ufeeb', '\ufeec', '\ufeea'}
	arWaw            = arabicHarf{'\u0648', '\ufeed', '\u0648', '\ufeee', '\ufeee'}
	arYehFarsi       = arabicHarf{'\u06cc', '\ufbfc', '\ufbfe', '\ufbff', '\ufbfd'}
	arYeh            = arabicHarf{'\u064a', '\ufef1', '\ufef3', '\ufef4', '\ufef2'}
	arAlefMaksura    = arabicHarf{'\u0649', '\ufeef', '\u0649', '\ufef0', '\ufef0'}
	arTatweel        = arabicHarf{'\u0640', '\u0640', '\u0640', '\u0640', '\u0640'}
	arLamAlef        = arabicHarf{'\ufefb', '\ufefb', '\ufefb', '\ufefc', '\ufefc'}
	arLamAlefHamzaA  = arabicHarf{'\ufef7', '\ufef7', '\ufef7', '\ufef8', '\ufef8'}
	arLamAlefHamzaB  = arabicHarf{'\ufef9', '\ufef9', '\ufef9', '\ufefa', '\ufefa'}
	arLamAlefMadda   = arabicHarf{'\ufef5', '\ufef5', '\ufef5', '\ufef6', '\ufef6'}
)

var arabicAlphabet = []arabicHarf{
	arAlefHamzaAbove, arAlef, arAlefMadda, arHamza, arWawHamza, arAlefHamzaBelow,
	arYehHamza, arBeh, arPeh, arTeh, arTehMarbuta, arTheh, arJeem, arTcheh,
	arHah, arKhah, arDal, arThal, arReh, arJeh, arZain, arSeen, arSheen,
	arSad, arDad, arTah, arZah, arAin, arGhain, arFeh, arQaf, arKaf, arKeheh,
	arGaf, arLam, arMeem, arNoon, arHeh, arWaw, arYehFarsi, arYeh, arAlefMaksura,
	arTatweel, arLamAlef, arLamAlefHamzaA, arLamAlefHamzaB, arLamAlefMadda,
}

// Letters that do not connect to the following letter (right-joining only).
var arabicRightJoinOnly = map[rune]bool{
	arAlefHamzaAbove.Unicode: true, arAlefMadda.Unicode: true, arAlef.Unicode: true,
	arHamza.Unicode: true, arWawHamza.Unicode: true, arAlefHamzaBelow.Unicode: true,
	arTehMarbuta.Unicode: true, arDal.Unicode: true, arThal.Unicode: true,
	arReh.Unicode: true, arZain.Unicode: true, arJeh.Unicode: true, arWaw.Unicode: true,
	arAlefMaksura.Unicode: true,
	arLamAlef.Unicode: true, arLamAlefHamzaA.Unicode: true,
	arLamAlefHamzaB.Unicode: true, arLamAlefMadda.Unicode: true,
}

var arabicHarfByRune map[rune]arabicHarf

func init() {
	arabicHarfByRune = make(map[rune]arabicHarf, len(arabicAlphabet)*5)
	for _, h := range arabicAlphabet {
		arabicHarfByRune[h.Unicode] = h
		arabicHarfByRune[h.Isolated] = h
		arabicHarfByRune[h.Beginning] = h
		arabicHarfByRune[h.Medium] = h
		arabicHarfByRune[h.Final] = h
	}
}

func (h arabicHarf) matches(r rune) bool {
	return r == h.Unicode || r == h.Isolated || r == h.Beginning || r == h.Medium || r == h.Final
}

func lookupArabicHarf(r rune) (arabicHarf, bool) {
	h, ok := arabicHarfByRune[r]
	return h, ok
}
