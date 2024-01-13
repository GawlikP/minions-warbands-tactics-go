package gameObjects

import (
  "minions-warbands-tactics/constants"
  "minions-warbands-tactics/textures"
)

type MinionType int

const (
  MRat MinionType = iota
  MFish
  MBaltie
  MThreedy
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
    Animation:    textures.InitAnimation(),
    Moving:       false,
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
    Animation:   textures.InitAnimation(), 
    Moving:       false,
  }
}

func InitBaltieMinion(x, y int) Minion {
 return Minion{
    Type:       MBaltie,
    Health:     1,
    MaxHealth:  1,
    Damage:     1,
    Speed:      150,
    USprite:    Sprite{},
    Xpos:       float64(x*constants.TILESIZE),
    Ypos:       float64(y*constants.TILESIZE),
    Defeated:     false,
    Direction:    4,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
    Animation:    textures.Animation{ CurrentAnimationFrame: 0,  Frames: 4 },
    Moving:       false,
  }
}

func InitThreedyMinion(x, y int) Minion {
 return Minion{
    Type:       MThreedy,
    Health:     1,
    MaxHealth:  1,
    Damage:     1,
    Speed:      200,
    USprite:    Sprite{},
    Xpos:       float64(x*constants.TILESIZE),
    Ypos:       float64(y*constants.TILESIZE),
    Defeated:     false,
    Direction:    4,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
    Animation:    textures.Animation{ CurrentAnimationFrame: 0,  Frames: 4 },
    Moving:       false,
  }
}
