package gameObjects

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/textures"
)

type MinionType int

const (
  MRat MinionType = iota
  MFish
)

type Minion struct {
  Type        MinionType
  Health      int64
  MaxHealth   int64
  Damage      int64
  Speed       int64
  USprite     Sprite
  Xpos        float64
  Ypos        float64
  Defeated    bool
}

func (u *Minion) Draw(screen *ebiten.Image, tex textures.Tex) {
  switch u.Type {
    case MRat:
      u.USprite.Draw(screen, tex.RatMinion)
    case MFish:
      u.USprite.Draw(screen, tex.FishMinion)
  }
}

func (u *Minion) Update() {
  u.USprite.Xpos = int(u.Xpos)
  u.USprite.Ypos = int(u.Ypos)
}
