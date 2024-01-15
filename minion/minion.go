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
  Type              MinionType
  Health            int64
  MaxHealth         int64
  Damage            int64
  Speed             int64
  USprite           gameObject.Sprite
  Xpos              float64
  Ypos              float64
  Defeated          bool
  Direction         byte
  Path              []int
  DestinationIndex  int
  PathIndex         int
  Animation         texture.Animation
  Moving            bool
  VX                float64
  VY                float64
  TargetIndex       int
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
  dx, dy := u.DestinationIndex%width, u.DestinationIndex/width
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
  u.DestinationIndex = -1
  return err
}

func (u *Minion) Update(grid []constant.BattleMapTileType, width int, ticks int) {
  u.HandleAnimation(ticks)
  if u.DestinationIndex != -1 {
    u.PathIndex = -1
    u.Path = []int{}
    u.TargetIndex = -1
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
