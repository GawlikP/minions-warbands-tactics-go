package minion

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/texture"
  "minions-warbands-tactics/pathfinder"
  "minions-warbands-tactics/constant"
  "minions-warbands-tactics/gameObject"
  "log"
)

type Minion struct {
  Type        MinionType
  Health      int64
  MaxHealth   int64
  Damage      int64
  Speed       int64
  USprite     gameObject.Sprite
  Xpos        float64
  Ypos        float64
  Defeated    bool
  Direction   byte
  Path        []int
  TargetIndex int
  PathIndex   int
  Animation   texture.Animation
  Moving      bool
  VX          float64
  VY          float64
}

func (u *Minion) Draw(screen *ebiten.Image, tex texture.Tex) {

  if u.Animation.Frames == 0 {
    u.DrawWithoutAnimations(screen, tex)
  } else {
    u.DrawPropperAnimation(screen, tex)
  }
}

func (u *Minion) GeneratePath(grid []constant.BattleMapTileType, width int) (error) {
  var err error
  //StartX, StartY, TargetX, TargetY, Width, Grid
  // POINT 0,0 dont work
  dx, dy := u.TargetIndex%width, u.TargetIndex/width
  log.Printf("PATHFINDING, DESTINATION DX:%d DY:%d", dx, dy)
  u.Path, err = pathfinder.AStar(
    int(u.Xpos/constant.TILESIZE),
    int(u.Ypos/constant.TILESIZE),
    dx,
    dy,
    width,
    grid, 
  )
  if err == nil {
    u.PathIndex = len(u.Path) - 1
    log.Printf("Path: %v", u.Path)
    for i, v := range u.Path  {
      log.Printf("Path[%d]: X:%d Y:%d", i, v%width, v/width)
    }
  }
  u.TargetIndex = -1
  return err
}

func (u *Minion) Update(grid []constant.BattleMapTileType, width int, ticks int) {
  u.HandleAnimation(ticks)
  if u.TargetIndex != -1 {
    u.PathIndex = -1
    u.Path = []int{}
    err := u.GeneratePath(grid, width)
    if err != nil {
      log.Printf("PATHFINDING ERROR: %v",  err)
    } else {
      log.Print("PathSet")
    }
  } else {
    u.MoveOnPath(width)
  }
  u.Move()
  u.USprite.Xpos = int(u.Xpos)
  u.USprite.Ypos = int(u.Ypos)
}

func (u *Minion) MoveOnPath(width int) {
  if u.PathIndex != -1 && len(u.Path) != 0 {
    x, y := u.Path[u.PathIndex]%width*constant.TILESIZE, u.Path[u.PathIndex]/width*constant.TILESIZE
    if int(u.Xpos) > x - 6 && int(u.Xpos) < x + 6 && int(u.Ypos) > y - 6 && int(u.Ypos) < y + 6 {
      u.PathIndex -= 1
    } else {
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
  } else {
    u.Path = []int{}
  }
}

func (u* Minion) Move() {
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
