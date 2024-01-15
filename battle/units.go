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
      b.Allies[idx].FollowEnemy(b.Enemies[b.Allies[idx].TargetIndex])
    }
    b.Allies[idx].Update(b.Tiles, b.Width, ticks)
  }
}

func (b *BattleMap) UpdateEnemies(ticks int) {
  for idx := range b.Enemies {
    // b.Enemies[idx].TargetEnemy(b.Allies, b.Tiles, b.Width)
    // if b.Enemies[idx].TargetIndex != -1 {
    //   b.Enemies[idx].FollowEnemy(b.Allies[b.Enemies[idx].TargetIndex])
    // }
    b.Enemies[idx].Update(b.Tiles, b.Width, ticks)
  }

  indexesToRemove := []int{}
  for i, v := range b.Enemies {
    if v.Health < 0 {
      indexesToRemove = append(indexesToRemove, i) 
    }
  }

  for idx := range indexesToRemove { 
    b.Effects = append(b.Effects[:idx], b.Effects[idx+1:]...)
  }
} 

func (b *BattleMap) RenderMinions(screen *ebiten.Image, tex texture.Tex) {
  for idx := range b.Allies {
    b.Allies[idx].Draw(screen, tex)
  }

  for idx := range b.Enemies {
    b.Enemies[idx].Draw(screen, tex)
  }
}

