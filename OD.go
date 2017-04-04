package main

import (
	"fmt"
	"image"
	"strings"

	"os"

	"image/draw"
	"image/png"

	"flag"

	"./GFX"
)

var hue = png.FormatError("HUE")

//var teststring = "abcdefghijklmnopqrstuvwxyz\n" + strings.ToUpper("abcdefghijklmnopqrstuvwxyz") + "\n" + "1234567890\n" + "~!@#$%^&*()_+{}|:\"<>?`-=[]\\;',./"

//var teststring = "Did someone say pancakes?"

//var face = "niko_pancakes"

var face string
var text string
var out string

func init() {
	flag.StringVar(&face, "face", "niko", "The filename of the face you want in the image")
	flag.StringVar(&text, "text", "SAMPLE", "Text to create")
	flag.StringVar(&out, "out", "out", "output file name")
	flag.Parse()
}

func main() {
	mainstring := text
	mainarray := strings.Split(mainstring, "\\n")
	fmt.Println("Starting drawing operation...\nWith face \"" + face + "\" and text:")
	for _, s := range mainarray {
		fmt.Println(s)
	}
	im := getimage()
	var img = image.NewRGBA(im.Bounds())
	draw.Draw(img, im.Bounds(), im, image.ZP, draw.Src)

	var let image.Image
	for {
		var needsrecalc bool
		for i, s := range mainarray {
			ok, usethis, more := calcstring(s)
			if !ok {
				mainarray[i] = usethis
				mainarray = append(mainarray[:i+1], append(more, mainarray[1+i:]...)...)
				needsrecalc = true
				break
			}
			for i2, r := range s {
				let = abc.Letter(r)
				sr := let.Bounds()
				pt := image.Pt(1*(letterleftmargin+(i2*10)), 1*(lettertopmargin+(i*21)))
				r := sr.Sub(sr.Min).Add(pt)
				draw.Draw(img, r, let, sr.Min, draw.Over)
			}
		}
		if !needsrecalc {
			break
		}
	}

	var faceimg image.Image

	data, err := os.Open("GFX/FS/" + face + ".png")
	if err == nil {
		faceimg, _, err = image.Decode(data)
		HE(err)
		sr := faceimg.Bounds()
		pt := image.Pt(img.Bounds().Dx()-(sr.Dx()+18), 18)
		r := sr.Sub(sr.Min).Add(pt)
		draw.Draw(img, r, faceimg, sr.Min, draw.Over)
	} else {
		fmt.Println("Couldnt load face:\n", err)
	}

	f, err := os.Create("output/" + out + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func getimage() image.Image {
	data, err := os.Open("GFX/TB.png")
	HE(err)
	img, _, err := image.Decode(data)
	HE(err)
	return img
}

func HE(err error) {
	if err != nil {
		panic(err)
	}
}

// could be 48, needs research
const maxletters = 47
const maxrows = 4
const letterleftmargin = 22
const lettertopmargin = 19

func calcstring(s string) (bool, string, []string) {
	if len(s) >= maxletters {
		totalwords := strings.Split(s, " ")
		var workstring string
		var firstretstring string
		var morestrings = []string{}
		var workingonfirst = true

		for i := 0; i < len(totalwords); i++ {
			if len(totalwords[i]) > maxletters {
				panic("WORD IS TOO LONG")
			}

			if len(workstring+" "+totalwords[i]) < maxletters || (workstring == "" && len(totalwords[i]) < maxletters) {
				if workstring == "" && len(totalwords[i]) < maxletters {
					workstring = totalwords[i]
				} else {
					workstring = workstring + " " + totalwords[i]
				}
			} else {
				if workingonfirst {
					firstretstring = workstring
					workingonfirst = false
				} else {
					morestrings = append(morestrings, workstring)
				}
				workstring = totalwords[i]
			}
		}
		morestrings = append(morestrings, workstring)
		return false, firstretstring, morestrings
	} else {
		return true, s, []string{""}
	}
}