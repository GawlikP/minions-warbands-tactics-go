package gameObjects

type BattleMapTileType int

const (
  Grass BattleMapTileType = iota
  Sand
  Stone
)

type BattleMap struct {
  Tiles   []BattleMapTileType
  Units   []Unit
  Width   int
}

func (b *BattleMap) Update() {}

func (b *BattleMap) Draw() {}

func (b *BattleMap) Input() {}

func (b *BattleMap) Init() {}
