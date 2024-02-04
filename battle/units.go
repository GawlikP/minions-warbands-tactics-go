package battle

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/texture"
)

func (b *BattleMap) UpdateAllies(ticks int) {
  for idx := range b.Allies {
    if b.Allies[idx].PathIndex == -1 && b.Allies[idx].TargetIndex == -1 {
      b.Allies[idx].TargetEnemy(b.Enemies, b.Tiles, b.Width)
    }
    if b.Allies[idx].TargetIndex != -1 {
      b.Allies[idx].FollowEnemy(&b.Enemies[b.Allies[idx].TargetIndex])
      effects := b.Allies[idx].Attack(&b.Enemies[b.Allies[idx].TargetIndex])
      if len(effects) > 0 {
        b.Effects = append(b.Effects, effects...)
      }
    }
    b.Allies[idx].Update(b.Tiles, b.Width, ticks)
  }
}

func (b *BattleMap) UpdateEnemies(ticks int) {
  for idx := range b.Enemies {
    if b.Enemies[idx].PathIndex == -1 && b.Enemies[idx].TargetIndex == -1 {
      b.Enemies[idx].TargetEnemy(b.Allies, b.Tiles, b.Width)
    }
    if b.Enemies[idx].TargetIndex != -1 {
      b.Enemies[idx].FollowEnemy(&b.Allies[b.Enemies[idx].TargetIndex])
      effects := b.Enemies[idx].Attack(&b.Allies[b.Enemies[idx].TargetIndex])
      if len(effects) > 0 {
        b.Effects = append(b.Effects, effects...)
      }
    }
    b.Enemies[idx].Update(b.Tiles, b.Width, ticks)
  }

  indexesToRemove := []int{}
  for i, v := range b.Enemies {
    if v.Health < 0 {
      indexesToRemove = append(indexesToRemove, i) 
    }
  }

  removed := 0
  for _, v := range indexesToRemove { 
    for idx := range b.Allies {
      if b.Allies[idx].TargetIndex == v {
        b.Allies[idx].TargetIndex = -1
      }
    }
    index := v - removed
    if index < 0 {
      index = 0
    }
    b.Enemies = append(b.Enemies[:index], b.Enemies[index+1:]...)
    removed++
  }
} 

func (b *BattleMap) RenderMinions(screen *ebiten.Image, tex texture.Tex) {
  for idx := range b.Enemies {
    b.Enemies[idx].Draw(screen, tex)
  }
  for idx := range b.Allies {
    b.Allies[idx].Draw(screen, tex)
  }
}

