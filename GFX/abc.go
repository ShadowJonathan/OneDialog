package abc

import "image"

func Letter(r rune) image.Image {

}

type letter struct {
	toppadding int
	data       [][8]bool
}

// default margin between letters: 2px
// letter width: 8px
// all letter's default top padding (cept "`") seems to be 2

const dp int = 2

var letters = map[rune]letter{
	'a': letter{
		toppadding: dp + 2,
		data:       [][8]bool{nosides, tworight, tworight, allbutoneleft, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneleft},
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
		data:       [][8]bool{allbutoneleft, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, sidestwo, allbutoneleft},
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
		data:       [][8]bool{[8]bool{false, false, false, false, false, true, true, false}, [8]bool{false, false, false, false, false, true, true, false}, empty, empty},
	},
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

var middletwo = [8]bool{false, false, false, true, true, false, false, false}
