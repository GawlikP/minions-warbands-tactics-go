package ui

import (
  "minions-warbands-tactics/textures"
  "github.com/hajimehoshi/ebiten/v2"
)

type InfoBadge struct {
  Xpos    int
  Ypos    int
  Msg     string
  Tag     string
  Active  bool
  sceneW  int
  sceneH  int

}

func (i *InfoBadge) Draw(screen *ebiten.Image, tex textures.Tex){
  if i.Active {
    i.RenderBadgeInTheCorner(screen, tex) 
  }
}

func (i *InfoBadge) Init(tag string, sceneW, sceneH int) {
  i.sceneW = sceneW
  i.sceneH = sceneH
  i.Xpos = sceneW - 240
  i.Ypos = 0 + 10
  i.Msg = ""
  i.Active = false
  i.Tag = tag
}

func (i* InfoBadge) RenderBadgeInTheCorner(screen *ebiten.Image, tex textures.Tex) {
  Op := &ebiten.DrawImageOptions{}
  Op.GeoM.Reset()
  Op.GeoM.Translate(float64(i.Xpos), float64(i.Ypos))
  screen.DrawImage(tex.UIBadge, Op)
  textures.DrawCenteredText(screen,tex.StandardFont, i.Msg, i.sceneW-i.Xpos/2, 32)
}
