package gameObject

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/texture"
  "minions-warbands-tactics/constant"
)

type BattleMapTileType int

const (
  Grass BattleMapTileType = iota
  Sand
  Stone
)

type BattleMap struct {
  Tiles   []BattleMapTileType
  Minions []Minion
  Width   int
}

func (b *BattleMap) Update(ticks int) {
  b.UpdateMinions(ticks)
}

func (b *BattleMap) Draw(screen *ebiten.Image, tex texture.Tex) {
  b.RenderTiles(screen, tex)
  b.RenderMinions(screen, tex)
}

func (b *BattleMap) Input() {}

func (b *BattleMap) Init() {}

func (b *BattleMap) RenderMinions(screen *ebiten.Image, tex texture.Tex) {
  for idx := range b.Minions {
    b.Minions[idx].Draw(screen, tex)
  }
}

func (b *BattleMap) RenderTiles(screen *ebiten.Image, tex texture.Tex) {
  Op := &ebiten.DrawImageOptions{}
  x, y := 0, 0

  for i, v := range b.Tiles {
    x = i%b.Width
    if x == 0 && i != 0 {
      y += 1
    }

    Op.GeoM.Reset()
    Op.GeoM.Translate(float64(x*constant.TILESIZE), float64(y*constant.TILESIZE))
    switch v {
      case Grass:
        screen.DrawImage(tex.GrassTile, Op)
      case Stone:
        screen.DrawImage(tex.StoneTile, Op)
      case Sand:
        screen.DrawImage(tex.SandTile, Op)
    }
  }
}

func (b *BattleMap) GetCurrentTileName(x, y int) string {
  index := b.GetTileIndex(x, y)
  if index == -1 {
    return "Out of Map"
  }
  switch b.Tiles[index] {
    case Grass:
      return "Grass"
    case Stone:
      return "Stone" 
    case Sand:
      return "Sand"
  }
  return ""
}

func (b *BattleMap) GetTileIndex(x, y int) int {
  index := x/constant.TILESIZE +  y/constant.TILESIZE * b.Width
  if index > len(b.Tiles)-1 || index < 0 {
    return -1
  }
  return index
}

func (b *BattleMap) UpdateMinions(ticks int) {
  for idx := range b.Minions {
    b.Minions[idx].Update(b.Tiles, b.Width, ticks)
  }
}
