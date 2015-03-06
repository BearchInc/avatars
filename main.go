package main

import (
	_ "code.google.com/p/draw2d/draw2d"
	"fmt"
	c2 "github.com/unseen/avatars/color"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	_ "math"
	"os"
)

type Names struct {
	Names []string
}

type Color struct {
	Name  string
	Value string
}

type Colors struct {
	Colors []Color
}

var names = []string{
	"star",
	"car",
	"outlet",
	"rocket",
	"clover",
	"strawberry",
	"coffee_pot",
	"rainbow",
	"lemon",
	"bicycle",
	"house",
	"paper_clip",
	"pizza",
	"coffee_cup",
	"carrot",
	"globe",
	"cloud_flake",
	"mouse",
	"lightbulb",
}

var colors = []Color{
	{"green", "#39B54A"},
	{"purple", "#7F47DD"},
	{"blue", "#22AFCA"},
	{"red", "#F15A24"},
	{"yellow", "#FBAE17"},
	{"pumpkin", "#F7931E"},
	{"forest", "#006837"},
	{"navy", "#0071BC"},
	{"teal", "#00A99D"},
	{"orange", "#F7931E"}, // orange is the new pumpkin
}

func findImage(name string) image.Image {
	reader, err := os.Open(fmt.Sprintf("images/%v.png", name))

	if err != nil {
		panic(err)
	}

	m, _, err := image.Decode(reader)

	if err != nil {
		panic(err)
	}

	return m
}

func createImageWithBackgroundColor(name string, c Color) {
	in := findImage(name)
	background := image.NewUniform(c2.Hex(c.Value))
	dst := image.NewRGBA(in.Bounds())

	// draw background
	draw.Draw(dst, dst.Bounds(), background, image.ZP, draw.Over)

	// draw picture of avatar
	draw.Draw(dst, dst.Bounds(), in, image.ZP, draw.Over)

	out, err := os.Create(fmt.Sprintf("out/%v_%v.png", c.Name, name))

	if err != nil {
		panic(err)
	}

	err = png.Encode(out, dst)

	if err != nil {
		panic(err)
	}
}

func createInverted(name string) {
	in := findImage(name)

	bounds := in.Bounds()

	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := in.At(x, y)
			r, g, b, a := c.RGBA()
			dst.Set(x, y, color.RGBA{uint8(65535 - r), uint8(65535 - g), uint8(65535 - b), uint8(65535 - a)})
		}
	}

	out, err := os.Create(fmt.Sprintf("out/inverted_%v.png", name))

	if err != nil {
		panic(err)
	}

	err = png.Encode(out, dst)

	if err != nil {
		panic(err)
	}
}

func main() {
	for _, name := range names {
		for _, color := range colors {
			createImageWithBackgroundColor(name, color)
		}
		createInverted(name)
	}

	createImageWithBackgroundColor("new", Color{"", "#008000"})
	createImageWithBackgroundColor("creator", Color{"", "#4BACC6"})
}
