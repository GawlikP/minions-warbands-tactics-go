package minion

import (
  "minions-warbands-tactics/constant"
  "minions-warbands-tactics/physics"
  "slices"
  "log"
)

func (u *Minion) MoveOnPath(width int) {
  if u.PathIndex != -1 && len(u.Path) != 0 {
    x, y := u.Path[u.PathIndex]%width*constant.TILESIZE, u.Path[u.PathIndex]/width*constant.TILESIZE
    if int(u.Xpos) > x - 6 && int(u.Xpos) < x + 6 && int(u.Ypos) > y - 6 && int(u.Ypos) < y + 6 {
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
  if u.VX == 0.0 && u.VY == 0.0 {
    u.Moving = false
  } else {
    u.Moving = true
  }
  u.Xpos += u.VX
  u.Ypos += u.VY
  u.VX = 0.0
  u.VY = 0.0
}


func (u *Minion) TargetEnemy(enemies []Minion, tiles []constant.BattleMapTileType, width int) {
  enemiesToSkip := u.EnemiesBehindWalls(enemies, tiles, width)
  if u.TargetIndex != -1 {
    return
  }
  for i, v := range enemies {
    if slices.Contains(enemiesToSkip, i) {
      continue
    }
    if physics.IsColidingOnCircle(
      int(v.Xpos-constant.UNITSIZE), 
      int(v.Ypos+constant.UNITSIZE),
      1,
      1,
      int(u.Xpos),
      int(u.Ypos),
      constant.ENEMYDETECTIONR,
    ) {
        log.Printf("NEW Target Index: %d", u.TargetIndex)
        u.TargetIndex = i
        break
    }
  }
}

func (u *Minion) FollowEnemy(enemy Minion) {
  u.MoveCloserTo(int(enemy.Xpos), int(enemy.Ypos))
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
      u.VX += float64(y) - u.Ypos
    }
  } else if u.Ypos > float64(y) {
    u.Direction = 4
    if possible_speed < u.Ypos - float64(y) {
      u.VY -= possible_speed
    } else {
      u.VX -= u.Ypos - float64(y)
    }
  }
}

func (u *Minion) EnemiesBehindWalls(enemies []Minion, tiles []constant.BattleMapTileType, width int) []int {
  enemiesIdx := []int{}
  for tidx, t := range tiles {
    if t != constant.Stone {
      continue
    }
    for idx, e := range enemies {
      // log.Printf("u x:%d y:%d e x:%d y:%d b x:%d y:%d", 
      // int(u.Xpos)/constant.TILESIZE, int(u.Ypos)/constant.TILESIZE, int(e.Xpos)/constant.TILESIZE, int(e.Ypos)/constant.TILESIZE, tidx%width, tidx/width)
      if physics.IsInLine(
        int(u.Xpos/constant.TILESIZE),
        int(u.Ypos/constant.TILESIZE),
        int(e.Xpos/constant.TILESIZE),
        int(e.Ypos/constant.TILESIZE),
        (tidx%width),
        (tidx/width),
      ) {
        enemiesIdx = append(enemiesIdx, idx)
      }
    }
  }
  return enemiesIdx
}
