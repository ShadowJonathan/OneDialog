package main

import (
	"fmt"
	"image"
	"io/ioutil"

	"image/png"

	"os"

	"image/color"

	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	data, err := ioutil.ReadFile("GFX/TerminusTTF-Bold.ttf")
	if err != nil {
		panic(err)
	}
	Font, err := freetype.ParseFont(data)
	if err != nil {
		panic(err)
	}

	size := 25

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

	ctx.DrawString("test", fixed.P(0, size))
	ctx.DrawString("abcdefghijklmnopqrstuvwxyz", fixed.P(0, size*2))
	ctx.DrawString(strings.ToUpper("abcdefghijklmnopqrstuvwxyz"), fixed.P(0, size*3))
	fmt.Println(saveimage(rgba.SubImage(rgba.Bounds())))
}

func saveimage(img image.Image) error {
	f, err := os.Create("out.png")
	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		return err
	}
	return nil
}
