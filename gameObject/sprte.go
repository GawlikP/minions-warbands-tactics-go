package gameObject

import (
  "github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
  Width       int
  Height      int
  Xpos        int
  Ypos        int
}

func (s *Sprite) Draw(screen *ebiten.Image, texture *ebiten.Image) {
  Op := &ebiten.DrawImageOptions{}
  Op.GeoM.Reset()
  Op.GeoM.Translate(float64(s.Xpos), float64(s.Ypos))
  screen.DrawImage(texture, Op)
}
