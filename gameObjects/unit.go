package gameObjects

type UnitType int

const (
  Worm UnitType = iota
)

type Unit struct {
  Type        UnitType
  Health      int64
  MaxHealt    int64
  Damage      int64
  Speed       int64
  UnitSprite  Sprite
}
