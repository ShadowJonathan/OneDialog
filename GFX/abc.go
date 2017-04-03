package abc

import "image"
import "image/draw"
import "image/color"

func Letter(r rune) image.Image {
	rgba := image.NewRGBA(image.Rect(0, 0, 10, 20))
	draw.Draw(rgba, rgba.Bounds(), image.Transparent, image.ZP, draw.Src)
	theletter, ok := letters[r]
	if !ok {
		return rgba
	}
	var y int
	y = theletter.toppadding
	for _, drawinfo := range theletter.data {
		for i, b := range drawinfo {
			if b {
				rgba.Set(i+1, y, color.White)
			}
		}
		y++
	}
	return rgba
}

type letter struct {
	toppadding int
	data       [][8]bool
}

// default margin between letters: 2px
// letter width: 8px
// all letter's default top padding (cept "`") seems to be 2

// ...
// note to self, make less boilerplate and a handy generate for this
// >.>

/*
àèìòùáéíóúñäëÿüïöâêîôûçæœ
¡™£¢∞§¶•ªº–≠∑®†¥øπ“‘«åß∂ƒ©˙∆˚¬…Ω≈√∫µ≤≥÷⁄€‹›ﬁﬂ‡°·‚—±„´‰ˇ¨ˆ∏»˝¸˛◊ı˜¯˘¿
*/

//var teststring = "abcdefghijklmnopqrstuvwxyz\n" + strings.ToUpper("abcdefghijklmnopqrstuvwxyz") + "\n" + "1234567890\n" + "~!@#$%^&*()_+{}|:\"<>?`-=[]\\;',./"

const dp int = 2

var letters = map[rune]letter{
	'a': letter{
		toppadding: dp + 4,
		data:       [][8]bool{nosides, tworight, tworight, allbutoneleft, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneleft},
	},
	'A': letter{
		toppadding: dp,
		data:       [][8]bool{nosides, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, full, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo},
	},
	'b': letter{
		toppadding: dp,
		data:       [][8]bool{twoleft, twoleft, twoleft, twoleft, allbutoneright, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneright},
	},
	'B': letter{
		toppadding: dp,
		data:       [][8]bool{allbutoneright, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneright, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneright},
	},
	'c': letter{
		toppadding: dp + 4,
		data:       [][8]bool{nosides, sidestwo, twoleft, twoleft, twoleft, twoleft, twoleft, sidestwo, nosides},
	},
	'C': letter{
		toppadding: dp,
		data:       [][8]bool{nosides, sidestwo, sidestwo, twoleft, twoleft, twoleft, twoleft, twoleft, twoleft, twoleft, sidestwo, sidestwo, nosides},
	},
	'd': letter{
		toppadding: dp,
		data:       [][8]bool{tworight, tworight, tworight, tworight, allbutoneleft, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneleft},
	},
	'D': letter{
		toppadding: dp,
		data:       [][8]bool{allbuttworight, [8]bool{true, true, false, false, false, true, true, false}, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, [8]bool{true, true, false, false, false, true, true, false}, allbuttworight},
	},
	'e': letter{
		toppadding: dp + 4,
		data:       [][8]bool{nosides, sidestwo, sidestwo, sidestwo, full, twoleft, twoleft, sidestwo, nosides},
	},
	'E': letter{
		toppadding: dp,
		data:       [][8]bool{full, twoleft, twoleft, twoleft, twoleft, twoleft, allbuttworight, twoleft, twoleft, twoleft, twoleft, twoleft, full},
	},
	'f': letter{
		toppadding: dp,
		data:       [][8]bool{[8]bool{false, false, false, false, true, true, true, true}, middletwo, middletwo, middletwo, nosides, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo},
	},
	'F': letter{
		toppadding: dp,
		data:       [][8]bool{full, twoleft, twoleft, twoleft, twoleft, twoleft, allbuttworight, twoleft, twoleft, twoleft, twoleft, twoleft, twoleft},
	},
	'g': letter{
		toppadding: dp + 4,
		data:       [][8]bool{allbutoneleft, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneleft, tworight, tworight, nosides},
	},
	'G': letter{
		toppadding: dp,
		data:       [][8]bool{nosides, sidestwo, sidestwo, twoleft, twoleft, twoleft, [8]bool{true, true, false, false, true, true, true, true}, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, nosides},
	},
	'h': letter{
		toppadding: dp,
		data:       [][8]bool{twoleft, twoleft, twoleft, twoleft, allbutoneright, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo},
	},
	'H': letter{
		toppadding: dp,
		data:       [][8]bool{sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, full, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo},
	},
	'i': letter{
		toppadding: dp,
		data:       [][8]bool{middletwo, middletwo, empty, empty, [8]bool{false, false, true, true, true, false, false, false}, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, nosidestwo},
	},
	'I': letter{
		toppadding: dp,
		data:       [][8]bool{nosidestwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, nosidestwo},
	},
	'j': letter{
		toppadding: dp,
		data:       [][8]bool{[8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, empty, empty, [8]bool{false, false, false, false, true, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, true, true, false, false, true, true, false}, [8]bool{false, true, true, false, false, true, true, false}, nosidestwo},
	},
	'J': letter{
		toppadding: dp,
		data:       [][8]bool{[8]bool{false, false, false, false, true, true, true, true}, tworightshiftedone, tworightshiftedone, tworightshiftedone, tworightshiftedone, tworightshiftedone, tworightshiftedone, tworightshiftedone, tworightshiftedone, tworightshiftedone, [8]bool{true, true, false, false, false, true, true, false}, [8]bool{true, true, false, false, false, true, true, false}, [8]bool{false, true, true, true, true, true, false, false}},
	}, // couldn't find this one anywhere O.o
	'k': letter{
		toppadding: dp,
		data:       [][8]bool{twoleft, twoleft, twoleft, twoleft, sidestwo, [8]bool{true, true, false, false, false, true, true, false}, [8]bool{true, true, false, false, true, true, false, false}, [8]bool{true, true, false, true, true, false, false, false}, [8]bool{true, true, true, true, false, false, false, false}, [8]bool{true, true, false, true, true, false, false, false}, [8]bool{true, true, false, false, true, true, false, false}, [8]bool{true, true, false, false, false, true, true, false}, sidestwo},
	},
	'K': letter{
		toppadding: dp,
		data:       [][8]bool{sidestwo, sidestwo, [8]bool{true, true, false, false, false, true, true, false}, [8]bool{true, true, false, false, true, true, false, false}, [8]bool{true, true, false, true, true, false, false, false}, [8]bool{true, true, true, true, false, false, false, false}, [8]bool{true, true, true, false, false, false, false, false}, [8]bool{true, true, true, true, false, false, false, false}, [8]bool{true, true, false, true, true, false, false, false}, [8]bool{true, true, false, false, true, true, false, false}, [8]bool{true, true, false, false, false, true, true, false}, sidestwo, sidestwo},
	},
	'l': letter{
		toppadding: dp,
		data:       [][8]bool{[8]bool{false, false, true, true, true, false, false, false}, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, middletwo, nosidestwo},
	},
	'L': letter{
		toppadding: dp,
		data:       append(cop(twoleft, 12), [][8]bool{full}...),
	},
	'm': letter{
		toppadding: dp + 4,
		data:       append([][8]bool{allbutoneright}, cop([8]bool{true, true, false, true, true, false, true, true}, 8)...),
	},
	'M': letter{
		dp,
		append([][8]bool{sides, sidestwo, [8]bool{true, true, true, false, false, true, true, true}, full, [8]bool{true, true, false, true, true, false, true, true}}, cop(sidestwo, 8)...),
	},
	'n': letter{
		dp + 4,
		append([][8]bool{allbutoneright}, cop(sidestwo, 8)...),
	},
	'N': letter{dp,
		app(
			cop(sidestwo, 4),
			[][8]bool{[8]bool{true, true, true, false, false, false, true, true}, [8]bool{true, true, true, true, false, false, true, true}, [8]bool{true, true, false, true, true, false, true, true}, [8]bool{true, true, false, false, true, true, true, true}, [8]bool{true, true, false, false, false, true, true, true}},
			cop(sidestwo, 4),
		)},
	'o':  letter{dp + 4, app([][8]bool{nosides}, cop(sidestwo, 7), [][8]bool{nosides})},
	'O':  letter{dp, app([][8]bool{nosides}, cop(sidestwo, 11), [][8]bool{nosides})},
	'p':  letter{dp + 4, app([][8]bool{allbutoneright}, cop(sidestwo, 7), [][8]bool{allbutoneright}, cop(twoleft, 3))},
	'P':  letter{dp, app([][8]bool{allbutoneright}, cop(sidestwo, 5), [][8]bool{allbutoneright}, cop(twoleft, 6))},
	'q':  letter{dp + 4, app([][8]bool{allbutoneleft}, cop(sidestwo, 7), [][8]bool{allbutoneleft}, cop(tworight, 3))},
	'Q':  letter{dp, app([][8]bool{nosides}, cop(sidestwo, 9), [][8]bool{[8]bool{true, true, false, true, true, false, true, true}}, [][8]bool{[8]bool{true, true, false, false, true, true, true, true}}, [][8]bool{nosides}, [][8]bool{tworight})},
	'r':  letter{dp + 4, app([][8]bool{[8]bool{true, true, false, true, true, true, true, true}, [8]bool{true, true, true, true, false, false, false, false}, [8]bool{true, true, true, false, false, false, false, false}}, cop(twoleft, 6))},
	'R':  letter{dp, app([][8]bool{allbutoneright}, cop(sidestwo, 5), [][8]bool{allbutoneright, [8]bool{true, true, true, true, false, false, false, false}, [8]bool{true, true, false, true, true, false, false, false}, [8]bool{true, true, false, false, true, true, false, false}, [8]bool{true, true, false, false, false, true, true, false}, sidestwo}, [][8]bool{sidestwo})},
	's':  letter{dp + 4, [][8]bool{nosides, sidestwo, twoleft, twoleft, nosides, tworight, tworight, sidestwo, nosides}},
	'S':  letter{dp, [][8]bool{nosides, sidestwo, sidestwo, twoleft, twoleft, twoleft, nosides, tworight, tworight, tworight, sidestwo, sidestwo, nosides}},
	't':  letter{dp, app(cop(middletwo, 4), [][8]bool{nosides}, cop(middletwo, 7), [][8]bool{[8]bool{false, false, false, false, true, true, true, true}})},
	'T':  letter{dp, app([][8]bool{full}, cop(middletwo, 12))},
	'u':  letter{dp + 4, app(cop(sidestwo, 8), m(allbutoneleft))},
	'U':  letter{dp, app(cop(sidestwo, 12), m(nosides))},
	'v':  letter{dp + 4, app(cop(sidestwo, 3), cop(twoonbothsides, 3), [][8]bool{nosidestwo}, cop(middletwo, 2))},
	'V':  letter{dp, app(cop(sidestwo, 5), cop(twoonbothsides, 4), cop(nosidestwo, 2), cop(middletwo, 2))},
	'w':  letter{dp + 4, app(cop(sidestwo, 3), cop([8]bool{true, true, false, true, true, false, true, true}, 5), [][8]bool{nosides})},
	'W':  letter{dp, app(cop(sidestwo, 8), [][8]bool{[8]bool{true, true, false, true, true, false, true, true}, full, [8]bool{true, true, true, false, false, true, true, true}, sidestwo, sides})},
	'x':  letter{dp + 4, app(cop(sidestwo, 2), m(twoonbothsides), m(nosidestwo), m(middletwo), m(nosidestwo), m(twoonbothsides), cop(sidestwo, 2))},
	'X':  letter{dp, app(cop(sidestwo, 2), m([8]bool{true, true, true, false, false, true, true, true}), m(twoonbothsides), m(nosides), m(nosidestwo), m(middletwo), m(nosidestwo), m(nosides), m(twoonbothsides), m([8]bool{true, true, true, false, false, true, true, true}), cop(sidestwo, 2))},
	'y':  letter{dp + 4, app(cop(sidestwo, 8), m(allbutoneleft), cop(tworight, 2), m(nosides))},
	'Y':  letter{dp, app(cop(sidestwo, 3), cop(twoonbothsides, 3), m(nosidestwo), cop(middletwo, 6))},
	'z':  letter{dp + 4, [][8]bool{full, tworight, tworightshiftedone, [8]bool{false, false, false, false, true, true, false, false}, middletwo, [8]bool{false, false, true, true, false, false, false, false}, twoleftshiftedone, twoleft, full}},
	'Z':  letter{dp, app(m(full), cop(tworight, 3), m(tworightshiftedone), m([8]bool{false, false, false, false, true, true, false, false}), m(middletwo), m([8]bool{false, false, true, true, false, false, false, false}), m(twoleftshiftedone), cop(twoleft, 3), m(full))},
	'.':  letter{dp + 11, app(cop(middletwo, 2))},
	',':  letter{dp + 10, app(cop(middletwo, 3), m([8]bool{false, false, true, true, false, false, false, false}))},
	'/':  letter{dp + 1, app(cop(tworightshiftedone, 2), cop([8]bool{false, false, false, false, true, true, false, false}, 2), cop(middletwo, 2), cop([8]bool{false, false, true, true, false, false, false, false}, 2), cop(twoleftshiftedone, 2), cop(twoleft, 2))},
	'[':  letter{dp, app([][8]bool{nosidestwo}, cop([8]bool{false, false, true, true, false, false, false, false}, 11), m(nosidestwo))},
	']':  letter{dp, app([][8]bool{nosidestwo}, cop([8]bool{false, false, false, false, true, true, false, false}, 11), m(nosidestwo))},
	'\'': letter{0, app(cop(middletwo, 4))},
	'`':  letter{1, [][8]bool{[8]bool{false, false, true, true, false, false, false, false}, middletwo}},
	'"':  letter{0, app(cop(twoonbothsides, 3))},
	//'@':
	'-': letter{dp + 8, [][8]bool{full}},
	//'#':
	//'%':
	' ': letter{0, [][8]bool{empty}},
	'!': letter{dp, app(cop(middletwo, 9), cop(empty, 2), cop(middletwo, 2))},
	'?': letter{dp, [][8]bool{nosidestwo, twoonbothsides, sidestwo, sidestwo, [8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, true, true, false, false}, middletwo, middletwo, middletwo, empty, empty, middletwo, middletwo}},
	':': letter{dp + 4, [][8]bool{middletwo, middletwo, empty, empty, empty, empty, middletwo, middletwo}},
	'0': letter{dp, [][8]bool{nosides, sidestwo, sidestwo, sidestwo, [8]bool{true, true, false, false, false, true, true, true}, [8]bool{true, true, false, false, true, true, true, true}, [8]bool{true, true, false, true, true, false, true, true}, [8]bool{true, true, true, true, false, false, true, true}, [8]bool{true, true, true, false, false, false, true, true}, sidestwo, sidestwo, sidestwo, nosides}},
	'1': letter{dp, app([][8]bool{middletwo, [8]bool{false, false, true, true, true, false, false, false}, [8]bool{false, true, true, true, true, false, false, false}}, cop(middletwo, 9), m(nosides))},
	'2': letter{dp, [][8]bool{
		nosides,
		sidestwo,
		sidestwo,
		sidestwo,
		tworight,
		tworight,
		tworightshiftedone,
		[8]bool{false, false, false, false, true, true, true, false},
		middletwo,
		[8]bool{false, false, true, true, false, false, false, false},
		twoleftshiftedone,
		twoleft,
		full}},
	'3': letter{dp, [][8]bool{nosides, sidestwo, sidestwo, tworight, tworight, tworight, [8]bool{false, false, true, true, true, true, true, false}, tworight, tworight, tworight, sidestwo, sidestwo, nosides}},
	'4': letter{dp, [][8]bool{
		tworight,
		[8]bool{false, false, false, false, false, true, true, true},
		[8]bool{false, false, false, false, true, true, true, true},
		[8]bool{false, false, false, true, true, false, true, true},
		[8]bool{false, false, true, true, false, false, true, true},
		[8]bool{false, true, true, false, false, false, true, true},
		sidestwo, sidestwo, sidestwo, full, tworight, tworight, tworight,
	}},
	'5': letter{dp, app(m(full), cop(twoleft, 4), m(allbutoneright), cop(tworight, 4), cop(sidestwo, 2), m(nosides))},
	'6': letter{dp, app(m([8]bool{false, false, true, true, true, true, true, false}), m(twoleftshiftedone), cop(twoleft, 3), m(allbutoneright), cop(sidestwo, 6), m(nosides))},
	'7': letter{dp, app(m(full), cop(sidestwo, 2), m(tworight),
		cop(tworightshiftedone, 2),
		cop([8]bool{false, false, false, false, true, true, true, false}, 2),
		cop(middletwo, 5),
	)},
	'8': letter{dp, app(m(nosides), cop(sidestwo, 5), m(nosides), cop(sidestwo, 5), m(nosides))},
	'9': letter{dp, app(m(nosides), cop(sidestwo, 6), m(nosides), cop(tworight, 3), m(tworightshiftedone), m([8]bool{false, true, true, true, true, true, false, false}))},
}

var sides = [8]bool{true, false, false, false, false, false, false, true}
var sidestwo = [8]bool{true, true, false, false, false, false, true, true}

var nosides = [8]bool{false, true, true, true, true, true, true, false}
var nosidestwo = [8]bool{false, false, true, true, true, true, false, false}

var full = [8]bool{true, true, true, true, true, true, true, true}
var empty = [8]bool{false, false, false, false, false, false, false, false}

var allbutoneright = [8]bool{true, true, true, true, true, true, true, false}
var allbutoneleft = [8]bool{false, true, true, true, true, true, true, true}
var allbuttworight = [8]bool{true, true, true, true, true, true, false, false}
var allbuttwoleft = [8]bool{false, false, true, true, true, true, true, true}
var twoleft = [8]bool{true, true, false, false, false, false, false, false}
var tworight = [8]bool{false, false, false, false, false, false, true, true}

var twoleftshiftedone = [8]bool{false, true, true, false, false, false, false, false}
var tworightshiftedone = [8]bool{false, false, false, false, false, true, true, false}

var twoonbothsides = [8]bool{false, true, true, false, false, true, true, false}

var middletwo = [8]bool{false, false, false, true, true, false, false, false}

func cop(boolray [8]bool, count int) [][8]bool {
	var tot [][8]bool
	for i := 0; i < count; i++ {
		tot = append(tot, boolray)
	}
	return tot
}

func app(stuff ...[][8]bool) [][8]bool {
	var out [][8]bool
	for _, r := range stuff {
		out = append(out, r...)
	}
	return out
}

func m(b [8]bool) [][8]bool {
	return [][8]bool{b}
}
