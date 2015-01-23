package main

import (
  "image"
  "image/draw"
  "fmt"
  "os"
  c2 "github.com/unseen/avatars/color"
  "image/png"
  _ "math"
  _ "code.google.com/p/draw2d/draw2d"
)

type Names struct {
  Names []string
}

type Color struct {
  Name string
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

func createImage(name string, c Color) {
  in := findImage(name)
  background := image.NewUniform(c2.Hex(c.Value))
  dst := image.NewRGBA(in.Bounds())

  //center := dst.Bounds().Size()
  //xc := float64(center.X) / 2
  //yc := float64(center.Y) / 2
  //r := float64(dst.Bounds().Dy())/2

  // draw background
  draw.Draw(dst, dst.Bounds(), background, image.ZP, draw.Over)

  //gc := draw2d.NewGraphicContext(dst)
  //gc.SetStrokeColor(c2.Hex(c.Value))
  //gc.SetFillColor(c2.Hex(c.Value))
  //gc.SetLineWidth(1)
  //gc.MoveTo(xc, yc)
  //gc.ArcTo(xc, yc, r, r, 0, 2*math.Pi)
  //gc.Fill()

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

func main() {
  for _, name := range names {
    for _, color := range colors {
      createImage(name, color)
    }
  }
}
