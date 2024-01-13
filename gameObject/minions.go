package gameObject

import (
  "minions-warbands-tactics/constant"
  "minions-warbands-tactics/texture"
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
    Xpos:         float64(x*constant.TILESIZE),
    Ypos:         float64(y*constant.TILESIZE),
    Defeated:     false,
    Direction:    0,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
    Animation:    texture.InitAnimation(),
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
    Xpos:       float64(x*constant.TILESIZE),
    Ypos:       float64(y*constant.TILESIZE),
    Defeated:     false,
    Direction:    0,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
    Animation:   texture.InitAnimation(), 
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
    Xpos:       float64(x*constant.TILESIZE),
    Ypos:       float64(y*constant.TILESIZE),
    Defeated:     false,
    Direction:    4,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
    Animation:    texture.Animation{ CurrentAnimationFrame: 0,  Frames: 4 },
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
    Xpos:       float64(x*constant.TILESIZE),
    Ypos:       float64(y*constant.TILESIZE),
    Defeated:     false,
    Direction:    4,
    Path:         []int{},
    TargetIndex:  -1,
    PathIndex:    -1,
    Animation:    texture.Animation{ CurrentAnimationFrame: 0,  Frames: 4 },
    Moving:       false,
  }
}
