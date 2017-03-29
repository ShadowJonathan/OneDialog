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
}

var sides = [8]bool{true, false, false, false, false, false, false, true}
var sidestwo = [8]bool{true, true, false, false, false, false, true, true}

var nosides = [8]bool{false, true, true, true, true, true, true, false}
var nosidestwo = [8]bool{false, false, true, true, true, true, false, false}

var full = [8]bool{true, true, true, true, true, true, true, true}
var empty = [8]bool{false, false, false, false, false, false, false, false}

var allbutoneright = [8]bool{true, true, true, true, true, true, true, false}
var allbutoneleft = [8]bool{false, true, true, true, true, true, true, true}
var twoleft = [8]bool{true, true, false, false, false, false, false, false}
var tworight = [8]bool{false, false, false, false, false, false, true, true}
