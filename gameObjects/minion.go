package gameObjects

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/textures"
  "log"
  "minions-warbands-tactics/constants"
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
  Direction   byte
  Path        []int
  TargetIndex int
  PathIndex   int
}

func (u *Minion) Draw(screen *ebiten.Image, tex textures.Tex) {
  switch u.Type {
    case MRat:
      u.USprite.Draw(screen, tex.RatMinion)
    case MFish:
      u.USprite.Draw(screen, tex.FishMinion)
  }
}

func (u *Minion) GeneratePath(grid []BattleMapTileType, width int) (error) {
  var err error
  //StartX, StartY, TargetX, TargetY, Width, Grid
  // POINT 0,0 dont work
  dx, dy := u.TargetIndex%width, u.TargetIndex/width
  log.Printf("PATHFINDING, DESTINATION DX:%d DY:%d", dx, dy)
  u.Path, err = AStar(
    int(u.Xpos/constants.TILESIZE),
    int(u.Ypos/constants.TILESIZE),
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

func (u *Minion) Update(grid []BattleMapTileType, width int) {
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
  u.USprite.Xpos = int(u.Xpos)
  u.USprite.Ypos = int(u.Ypos)
}

func (u *Minion) MoveOnPath(width int) {
  if u.PathIndex != -1 && len(u.Path) != 0 {
    x, y := float64(u.Path[u.PathIndex]%width)*constants.TILESIZE, float64(u.Path[u.PathIndex]/width)*constants.TILESIZE
    if u.Xpos > x - 0.25 && u.Xpos < x + 0.25 && u.Ypos > y - 0.25 && u.Ypos < y + 0.25 {
      u.PathIndex -= 1
    } else {
      if u.Xpos < x {
        u.Xpos += float64(u.Speed) / 100
      } else if u.Xpos > x {
        u.Xpos -= float64(u.Speed) / 100
      } else if u.Ypos < y {
        u.Ypos += float64(u.Speed) / 100
      } else if u.Ypos > y {
        u.Ypos -= float64(u.Speed) / 100
      }
    }
  } else {
    u.Path = []int{}
  }
}
