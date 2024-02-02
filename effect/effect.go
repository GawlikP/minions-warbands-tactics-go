package effect

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/texture"
  "minions-warbands-tactics/constant"
)

type Effect struct {
  Xpos, Ypos  int
  LifeTime    int
  CurrentTime int
  Animation   texture.Animation
  Type        constant.EffectType
}

func CreateEffect(x, y, l int, a texture.Animation, t constant.EffectType) Effect {
  return Effect{
    Xpos: x,
    Ypos: y,
    LifeTime: l,
    Animation: a,
    Type: t,
    CurrentTime: 0,
  }
}

func (e *Effect) Draw(screen *ebiten.Image, tex texture.Tex) {
  Op := &ebiten.DrawImageOptions{}
  Op.GeoM.Reset()
  Op.GeoM.Translate(float64(e.Xpos), float64(e.Ypos))
  switch e.Type {
    case constant.StandardParticle:
      screen.DrawImage(tex.StandardParticle[e.Animation.CurrentAnimationFrame], Op)
    case constant.TargetParticle:
      screen.DrawImage(tex.TargetParticle[e.Animation.CurrentAnimationFrame], Op)
  }
}

func (e *Effect) Update(ticks int) {
  if e.Animation.Frames != 0 {
    e.Animation.UpdateAnimationIndex(ticks)
  }
  e.CurrentTime += 1
}
