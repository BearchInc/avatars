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
  "avatar_alarm_clock",
  "avatar_apple",
  "avatar_battery",
  "avatar_bicycle",
  "avatar_blocks",
  "avatar_book",
  "avatar_calculator",
  "avatar_car",
  "avatar_carrot",
  "avatar_chain_link",
  "avatar_cheese",
  "avatar_chef_hat",
  "avatar_cigarette",
  "avatar_cloud_flake",
  "avatar_clover",
  "avatar_coffee_cup",
  "avatar_coffee_pot",
  "avatar_comb",
  "avatar_crosshair",
  "avatar_drumstick",
  "avatar_egg_cup",
  "avatar_eye",
  "avatar_fan",
  "avatar_film",
  "avatar_fries",
  "avatar_globe",
  "avatar_grape",
  "avatar_hanger",
  "avatar_heart",
  "avatar_house",
  "avatar_key",
  "avatar_lemon",
  "avatar_lightbulb",
  "avatar_loaf",
  "avatar_moon",
  "avatar_mouse",
  "avatar_musical_note",
  "avatar_outlet",
  "avatar_padlock",
  "avatar_paper_clip",
  "avatar_peapod",
  "avatar_pen_tip",
  "avatar_pencil",
  "avatar_phone",
  "avatar_pizza",
  "avatar_platter",
  "avatar_quote",
  "avatar_rainbow",
  "avatar_rocket",
  "avatar_shopping_cart",
  "avatar_silverware",
  "avatar_smoke_plume",
  "avatar_star",
  "avatar_storm",
  "avatar_strawberry",
  "avatar_suitcase",
  "avatar_sun",
  "avatar_tag",
  "avatar_tangerine",
  "avatar_teacup",
  "avatar_toaster",
  "avatar_tornado",
  "avatar_umbrella",
  "avatar_walkie_talkie",
  "avatar_water_drop",
  "avatar_webcam",
  "avatar_wine_glass",
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

  out, err := os.Create(fmt.Sprintf("out/inverted_%v%v.png", name, suffix))

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
    createInverted(name, "")
    createInverted(name, "@2x")
    createInverted(name, "@3x")
  }
}
