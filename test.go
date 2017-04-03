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
	/*data, err := ioutil.ReadFile("GFX/TerminusTTF-Bold.ttf")
	if err != nil {
		panic(err)
	}
	Font, err := freetype.ParseFont(data)
	if err != nil {
		panic(err)
	}

	size := 21

	ctx := freetype.NewContext()
	ctx.SetFont(Font)
	ctx.SetFontSize(float64(size))

	rgba := image.NewRGBA(image.Rect(0, 0, 2000, 200))
	for x := 0; x < 2000; x++ {
		for y := 0; y < 200; y++ {
			c := color.RGBA{0, 0, 0,
				255,
			}
			rgba.Set(x, y, c)
		}
	}
	ctx.SetDst(rgba)
	ctx.SetHinting(font.HintingNone)
	ctx.SetClip(rgba.Bounds())
	ctx.SetSrc(image.NewUniform(color.White))

	ctx.DrawString("1234567890", fixed.P(0, size))
	ctx.DrawString("abcdefghijklmnopqrstuvwxyz", fixed.P(0, size*2))
	ctx.DrawString(strings.ToUpper("abcdefghijklmnopqrstuvwxyz"), fixed.P(0, size*3))
	*/
	mainstring := text
	mainarray := strings.Split(mainstring, "\n")
	im := getimage()
	var img = image.NewRGBA(im.Bounds())
	draw.Draw(img, im.Bounds(), im, image.ZP, draw.Src)

	fmt.Println("beginning operation,", mainarray)
	var let image.Image
	for i, s := range mainarray {
		fmt.Println("making new sting...\n" + s)
		for i2, r := range s {
			let = abc.Letter(r)
			fmt.Println("drawing at", letterleftmargin+(i2*10), lettertopmargin+(i*21))
			sr := let.Bounds()
			pt := image.Pt(1*(letterleftmargin+(i2*10)), 1*(lettertopmargin+(i*21)))
			r := sr.Sub(sr.Min).Add(pt)
			draw.Draw(img, r, let, sr.Min, draw.Over)
		}
	}

	var faceimg image.Image

	data, err := os.Open("GFX/FS/" + face + ".png")
	if err == nil {
		faceimg, _, err = image.Decode(data)
		HE(err)
		//var faceimg = image.NewRGBA(fi.Bounds())
		//draw.Draw(faceimg, fi.Bounds(), fi, image.ZP, draw.Src)
		sr := faceimg.Bounds()
		pt := image.Pt(img.Bounds().Dx()-(sr.Dx()+18), 18)
		r := sr.Sub(sr.Min).Add(pt)
		fmt.Println(r)
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
