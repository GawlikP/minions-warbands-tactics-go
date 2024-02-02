package minion

import (
  "minions-warbands-tactics/constant"
  "minions-warbands-tactics/physics"
  "minions-warbands-tactics/effect"
  "minions-warbands-tactics/texture"
  "log"
  "math"
)

func (u *Minion) MoveOnPath(width int) {
  if u.PathIndex != -1 && len(u.Path) != 0 {
    x, y := u.Path[u.PathIndex]%width*constant.TILESIZE, u.Path[u.PathIndex]/width*constant.TILESIZE
    if int(math.Round(u.Xpos)) == x && int(math.Round(u.Ypos)) == y {
      u.PathIndex -= 1
    } else {
      u.MoveCloserTo(x, y)
    }
  } else {
    u.DestinationIndex = -1
    u.Path = []int{}
  }
}

func (u *Minion) Move() {
  if u.State != MSFighing {
    if u.VX == 0.0 && u.VY == 0.0 {
      u.State = MSIdle
    } else  {
      u.State = MSMoving
    }
  }
  u.Xpos += u.VX
  u.Ypos += u.VY
  u.VX = 0.0
  u.VY = 0.0
}


func (u *Minion) TargetEnemy(enemies []Minion, tiles []constant.BattleMapTileType, width int) {
  if u.TargetIndex != -1 {
    return
  }
  for i, v := range enemies {
    if !u.IsEnemyVisible(v, tiles, width) {
      continue
    }
    if physics.IsColidingOnCircle(
      int(v.Xpos-constant.UNITSIZE/2), 
      int(v.Ypos+constant.UNITSIZE/2),
      1,
      1,
      int(u.Xpos-constant.UNITSIZE/2),
      int(u.Ypos+constant.UNITSIZE/2),
      constant.ENEMYDETECTIONR,
    ) {
        log.Printf("NEW Target Index: %d", u.TargetIndex)
        u.TargetIndex = i
        break
    }
  }
}

func (u *Minion) FollowEnemy(e *Minion) {
  if int(math.Round(u.Xpos)/constant.TILESIZE) != int(math.Round(e.Xpos)/constant.TILESIZE) || int(math.Round(u.Ypos)/constant.TILESIZE) != int(math.Round(e.Ypos)/constant.TILESIZE) {
    u.MoveCloserTo(int(e.Xpos), int(e.Ypos))
  }
  if physics.IsColidingOnCircle( 
      int(e.Xpos-constant.UNITSIZE/2), 
      int(e.Ypos+constant.UNITSIZE/2),
      constant.UNITSIZE,
      constant.UNITSIZE,
      int(u.Xpos-constant.UNITSIZE/2),
      int(u.Ypos+constant.UNITSIZE/2),
      u.AttackRange,
  ) {
    u.State = MSFighing
  }
}

func (u *Minion) MoveCloserTo(x,y int) {
  possible_speed := float64(u.Speed) / 100
  if u.Xpos < float64(x) {
    u.Direction = 1
    if possible_speed < float64(x) - u.Xpos {
      u.VX += possible_speed
    } else {
      u.VX += float64(x) - u.Xpos
    }
  } else if u.Xpos > float64(x) {
    u.Direction = 2
    if possible_speed < u.Xpos - float64(x) {
      u.VX -= possible_speed
    } else {
      u.VX -= u.Xpos - float64(x)
    }
  } else if u.Ypos < float64(y) {
    u.Direction = 3
    if possible_speed < float64(y) - u.Ypos {
      u.VY += possible_speed
    } else {
      u.VY += float64(y) - u.Ypos
    }
  } else if u.Ypos > float64(y) {
    u.Direction = 4
    if possible_speed < u.Ypos - float64(y) {
      u.VY -= possible_speed
    } else {
      u.VY -= u.Ypos - float64(y)
    }
  }
}

func (u *Minion) IsEnemyVisible(enemy Minion, tiles []constant.BattleMapTileType, width int) bool {
  for tidx, t := range tiles {
    if t != constant.Stone {
      continue
    }
      // log.Printf("u x:%d y:%d e x:%d y:%d b x:%d y:%d", 
      // int(u.Xpos)/constant.TILESIZE, int(u.Ypos)/constant.TILESIZE, int(e.Xpos)/constant.TILESIZE, int(e.Ypos)/constant.TILESIZE, tidx%width, tidx/width)
    if physics.IsInLine(
      u.Xpos/constant.TILESIZE,
      u.Ypos/constant.TILESIZE,
      enemy.Xpos/constant.TILESIZE,
      enemy.Ypos/constant.TILESIZE,
      float64(tidx%width)-0.5,
      float64(tidx/width)-0.5,
    ) {
      return false
    }
  }
  // log.Printf("Indexes of enemies behind walls: %v", enemiesIdx)
  return true 
}
func (u *Minion) Attack(enemy *Minion) []effect.Effect {
  effects := []effect.Effect{}
  if u.TargetIndex != -1 {
    if u.PerformAttack == true {
      u.PerformAttack = false
      u.AttackCounter = 0
        enemy.Health -= u.Damage
        if enemy.Health < 0 {
          u.State = MSIdle
        }
        effects = append(effects, effect.Effect{
            Xpos: int(enemy.Xpos),
            Ypos: int(enemy.Ypos),
            Animation: texture.Animation{ CurrentAnimationFrame: 0, Frames: 6 },
            LifeTime: 60,
            Type: constant.StandardParticle,
            CurrentTime: 0,
        },
      )
      log.Print("Attacked!")
    }
  }
  return effects
}

func (u *Minion) TryToHit() {
  if u.TargetIndex != -1 {
    if u.Animation.CurrentAnimationFrame == u.AAttackIndex && u.PerformAttack != true && u.AttackCounter > u.AttackSpeed {
      u.PerformAttack = true
    }
  }
}
