package gameObjects

import "minions-warbands-tactics/constants"

type MinionType int

const (
  MRat MinionType = iota
  MFish
)

func InitRatMinion(x, y int) Minion {
 return Minion{
    Type:         MRat,
    Health:       1,
    MaxHealth:    1,
    Damage:       1,
    Speed:        100,
    USprite:      Sprite{},
    Xpos:         float64(x*constants.TILESIZE),
    Ypos:         float64(y*constants.TILESIZE),
    Defeated:     false,
    Direction:    0,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
  }
}

func InitFishMinion(x, y int) Minion {
 return Minion{
    Type:       MFish,
    Health:     1,
    MaxHealth:  1,
    Damage:     1,
    Speed:      1,
    USprite:    Sprite{},
    Xpos:       float64(x*constants.TILESIZE),
    Ypos:       float64(y*constants.TILESIZE),
    Defeated:     false,
    Direction:    0,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
  }
}
