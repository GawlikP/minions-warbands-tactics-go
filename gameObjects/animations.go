package gameObjects

import (
  "minions-warbands-tactics/textures"
  "github.com/hajimehoshi/ebiten/v2"
)

func (u *Minion) DrawWithoutAnimations(screen *ebiten.Image, tex textures.Tex) {
  switch u.Type {
    case MRat:
      u.USprite.Draw(screen, tex.RatMinion)
    case MFish:
      u.USprite.Draw(screen, tex.FishMinion)
  }
}

func (u *Minion) HandleAnimation(ticks int) {
  if u.Animation.Frames == 0  || u.Moving == false {
    return
  }
  u.Animation.UpdateAnimationIndex(ticks, 4)
}

func (u *Minion) DrawPropperAnimation(screen *ebiten.Image, tex textures.Tex) {
  switch u.Type {
  case MBaltie:
    if u.Direction == 1 {
      u.USprite.Draw(screen, tex.BazaltieWalkingRight[u.Animation.CurrentAnimationFrame])
    } else if u.Direction == 2 {
      u.USprite.Draw(screen, tex.BazaltieWalkingLeft[u.Animation.CurrentAnimationFrame])
    } else if u.Direction == 3 {
      u.USprite.Draw(screen, tex.BazaltieWalkingDown[u.Animation.CurrentAnimationFrame])
    } else if u.Direction == 4 {
      u.USprite.Draw(screen, tex.BazaltieWalkingUp[u.Animation.CurrentAnimationFrame])
    }
  case MThreedy:
    if u.Direction == 1 {
      u.USprite.Draw(screen, tex.ThreedyWalkingRight[u.Animation.CurrentAnimationFrame])
    } else if u.Direction == 2 {
      u.USprite.Draw(screen, tex.ThreedyWalkingLeft[u.Animation.CurrentAnimationFrame])
    } else if u.Direction == 3 {
      u.USprite.Draw(screen, tex.ThreedyWalkingDown[u.Animation.CurrentAnimationFrame])
    } else if u.Direction == 4 {
      u.USprite.Draw(screen, tex.ThreedyWalkingUp[u.Animation.CurrentAnimationFrame])
    }
  }
}
