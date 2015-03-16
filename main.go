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
  "alarm_clock",
  "apple",
  "battery",
  "bicycle",
  "blocks",
  "book",
  "calculator",
  "car",
  "carrot",
  "chain_link",
  "cheese",
  "chef_hat",
  "cigarette",
  "cloud_flake",
  "clover",
  "coffee_cup",
  "coffee_pot",
  "comb",
  "crosshair",
  "drumstick",
  "egg_cup",
  "eye",
  "fan",
  "film",
  "fries",
  "globe",
  "grape",
  "hanger",
  "heart",
  "house",
  "key",
  "lemon",
  "lightbulb",
  "loaf",
  "moon",
  "mouse",
  "musical_note",
  "outlet",
  "padlock",
  "paper_clip",
  "peapod",
  "pen_tip",
  "pencil",
  "phone",
  "pizza",
  "platter",
  "quote",
  "rainbow",
  "rocket",
  "shopping_cart",
  "silverware",
  "smoke_plume",
  "star",
  "storm",
  "strawberry",
  "suitcase",
  "sun",
  "tag",
  "tangerine",
  "teacup",
  "toaster",
  "tornado",
  "umbrella",
  "walkie_talkie",
  "water_drop",
  "webcam",
  "wine_glass",
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

func findImage(name string, suffix string) image.Image {
  reader, err := os.Open(fmt.Sprintf("images/%v%v.png", name, suffix))

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
  in := findImage(name, "")
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

func createInverted(name string, suffix string) {
  in := findImage(name, suffix)

  bounds := in.Bounds()

  dst := image.NewRGBA(bounds)

  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      c := in.At(x, y)
      r, g, b, a := c.RGBA()
      dst.Set(x, y, color.RGBA{uint8(65535 - r), uint8(65535 - g), uint8(65535 - b), uint8(65535 - a)})
    }
  }

  out, err := os.Create(fmt.Sprintf("out/inverted_%v.png", name ))

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
    //for _, color := range colors {
    //createImageWithBackgroundColor(name, color)
    //}
    createInverted(name, "@3x")
  }
}
