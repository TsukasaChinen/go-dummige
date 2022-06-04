package actions

import (
	"dummige/structs"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"unicode/utf8"

	"github.com/goki/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func CreateImage(i *structs.Images) error {

	c := i.C
	n := i.N
	w := i.W
	h := i.H

	var r string
	var g string
	var b string

	// Convert RGB by Color Code

	if utf8.RuneCountInString(c) == 3 {
		r = c[0:1] + c[0:1]
		g = c[1:2] + c[1:2]
		b = c[2:3] + c[2:3]
	} else {
		r = c[0:2]
		g = c[2:4]
		b = c[4:6]
	}

	rr, _ := strconv.ParseInt(r, 16, 64)
	gg, _ := strconv.ParseInt(g, 16, 64)
	bb, _ := strconv.ParseInt(b, 16, 64)

	rgb := new(structs.RGB)
	rgb.R = uint8(rr)
	rgb.G = uint8(gg)
	rgb.B = uint8(bb)

	// Init Image
	s := image.Point{0, 0}
	e := image.Point{w, h}

	img := image.NewRGBA(image.Rectangle{s, e})

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{rgb.R, rgb.G, rgb.B, 0xff})
		}
	}

	if w >= 40 && h >= 40 {
		err := CreateFont(i, rgb, img)
		if err != nil {
			return err
		}
	}

	// Create Image
	f, err := os.Create(n)
	png.Encode(f, img)

	return err
}

func CreateFont(i *structs.Images, rgb *structs.RGB, img *image.RGBA) error {

	c := i.C
	w := i.W
	h := i.H

	r := rgb.R
	g := rgb.G
	b := rgb.B

	s := image.Point{0, 0}
	e := image.Point{w, h}

	// Font Load
	ftBinary, err := ioutil.ReadFile("assets/MPLUSRounded1c-Regular.ttf")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	ft, err := truetype.Parse(ftBinary)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Font Size and Position
	var fs float64 = 10
	var fp int = 4

	if w >= 50 {
		fs = 11
		fp = 6
	}
	if w >= 100 {
		fs = 18
		fp = 8
	}
	if w >= 500 {
		fs = 24
		fp = 10
	}
	if w >= 800 {
		fs = 30
		fp = 15
	}

	// Init Font
	opt := truetype.Options {
		Size:              fs,
		DPI:               0,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	face := truetype.NewFace(ft, &opt)

	// Font Color
	max := uint8(200)
	fc := image.NewUniform(color.RGBA{255, 255, 255, 255})
	if r >= max && g >= max && b >= max {
		fc = image.NewUniform(color.RGBA{0, 0, 0, 255})
	}

	// Font Deawer
	d := &font.Drawer {
		Dst: img,
		Src: fc,
		Face: face,
		Dot: fixed.P((s.X + fp), (e.Y - (int(opt.Size) + (fp + 2)))),
	}

	if w > 40 && h > 40 {
		d.DrawString(strconv.Itoa(w) + "x" + strconv.Itoa(h))
		d.Dot = fixed.P((s.X + fp), (e.Y - fp))
		d.DrawString("#" + c)
	}

	return err
}

