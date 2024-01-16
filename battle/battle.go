package battle

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/texture"
  "minions-warbands-tactics/constant"
  "minions-warbands-tactics/minion"
  "minions-warbands-tactics/effect"
)

type BattleMap struct {
  Tiles   []constant.BattleMapTileType
  // Minions []minion.Minion
  Effects []effect.Effect
  Width   int
  Allies  []minion.Minion
  Enemies []minion.Minion
}

func (b *BattleMap) Update(ticks int) {
  b.UpdateEnemies(ticks)
  b.UpdateAllies(ticks)
  b.UpdateEffects(ticks)
}

func (b *BattleMap) Draw(screen *ebiten.Image, tex texture.Tex) {
  b.RenderTiles(screen, tex)
  b.RenderMinions(screen, tex)
  b.RenderEffects(screen, tex)
}

func (b *BattleMap) Input() {}

func (b *BattleMap) Init() {}

func (b *BattleMap) RenderEffects(screen *ebiten.Image, tex texture.Tex) {
  for idx := range b.Effects {
    b.Effects[idx].Draw(screen, tex)
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
      case constant.Grass:
        screen.DrawImage(tex.GrassTile, Op)
      case constant.Stone:
        screen.DrawImage(tex.StoneTile, Op)
      case constant.Sand:
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
    case constant.Grass:
      return "Grass"
    case constant.Stone:
      return "Stone" 
    case constant.Sand:
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

func (b *BattleMap) UpdateEffects(ticks int) {
  indexesToRemove := []int{}
  for idx := range b.Effects {
    b.Effects[idx].Update(ticks)
  }
  for idx := range b.Effects {
    if b.Effects[idx].LifeTime < b.Effects[idx].CurrentTime {
      indexesToRemove = append(indexesToRemove, idx)  
    }
  }
  removed := 0
  for _, v := range indexesToRemove {
    index := v - removed
    if index < 0 {
      index = 0
    }
    b.Effects = append(b.Effects[:index], b.Effects[index+1:]...)
    removed++
  }
}

func (b *BattleMap) AddEffect(e effect.Effect) {
  b.Effects = append(b.Effects, e)
}


