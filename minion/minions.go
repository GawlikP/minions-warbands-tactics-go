package minion

import (
  "minions-warbands-tactics/constant"
  "minions-warbands-tactics/texture"
  "minions-warbands-tactics/gameObject"
)

type MinionType int

const (
  MRat MinionType = iota
  MFish
  MBaltie
  MThreedy
)

type MinionState byte

const (
  MSIdle MinionState = iota
  MSMoving
  MSFighing
)

func InitRatMinion(x, y int) Minion {
 return Minion{
    Type:               MRat,
    Health:             20,
    MaxHealth:          1,
    Damage:             1,
    Speed:              1,
    USprite:            gameObject.Sprite{},
    Xpos:               float64(x*constant.TILESIZE),
    Ypos:               float64(y*constant.TILESIZE),
    Defeated:           false,
    Direction:          0,
    Path:               []int{},
    DestinationIndex:   -1,
    PathIndex:          -1,
    Animation:          texture.InitAnimation(),
    State:              MSIdle,
    TargetIndex:        -1,
    // ATTACKING
    PerformAttack:      false,
    AAttackIndex:       0,
    AttackSpeed:        40,
    AttackCounter:      0,
    AttackRange:        32,
  }
}

func InitFishMinion(x, y int) Minion {
 return Minion{
    Type:               MFish,
    Health:             1,
    MaxHealth:          1,
    Damage:             1,
    Speed:              1,
    USprite:            gameObject.Sprite{},
    Xpos:               float64(x*constant.TILESIZE),
    Ypos:               float64(y*constant.TILESIZE),
    Defeated:           false,
    Direction:          0,
    Path:               []int{},
    DestinationIndex:   -1,
    PathIndex:          -1,
    Animation:          texture.InitAnimation(), 
    State:              MSIdle,
    TargetIndex:        -1,
    // ATTACKING
    PerformAttack:      false,
    AAttackIndex:         3,
    AttackSpeed:        40,
    AttackCounter:      0,
    AttackRange:        32,
  }
}

func InitBaltieMinion(x, y int) Minion {
 return Minion{
    Type:               MBaltie,
    Health:             1,
    MaxHealth:          1,
    Damage:             1,
    Speed:              100,
    USprite:            gameObject.Sprite{},
    Xpos:               float64(x*constant.TILESIZE),
    Ypos:               float64(y*constant.TILESIZE),
    Defeated:           false,
    Direction:          4,
    Path:               []int{},
    DestinationIndex:   -1,
    PathIndex:          -1,
    Animation:          texture.Animation{ CurrentAnimationFrame: 0,  Frames: 4 },
    State:              MSIdle,
    TargetIndex:        -1,
    // ATTACKING
    PerformAttack:      false,
    AAttackIndex:       3,
    AttackSpeed:        600,
    AttackCounter:      0,
    AttackRange:        32,
  }
}

func InitThreedyMinion(x, y int) Minion {
 return Minion{
    Type:               MThreedy,
    Health:             1,
    MaxHealth:          1,
    Damage:             1,
    Speed:              200,
    USprite:            gameObject.Sprite{},
    Xpos:               float64(x*constant.TILESIZE),
    Ypos:               float64(y*constant.TILESIZE),
    Defeated:           false,
    Direction:          4,
    Path:               []int{},
    DestinationIndex:   -1,
    PathIndex:          -1,
    Animation:          texture.Animation{ CurrentAnimationFrame: 0,  Frames: 4 },
    State:              MSIdle,
    TargetIndex:        -1,
    // ATTACKING
    PerformAttack:      false,
    AAttackIndex:         3,
    AttackSpeed:        40,
    AttackCounter:      0,
    AttackRange:        32,
  }
}
