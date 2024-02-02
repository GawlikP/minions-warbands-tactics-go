package scene

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "minions-warbands-tactics/gameObject"
  "minions-warbands-tactics/texture"
  "minions-warbands-tactics/tilemap"
  "minions-warbands-tactics/ui"
  "minions-warbands-tactics/battle"
  "minions-warbands-tactics/minion"
  "minions-warbands-tactics/effect"
  "minions-warbands-tactics/constant"
  "log"
)

type BattleSimulationScene struct {
  State               SceneState
  sceneW              int
  sceneH              int
  Cursor              gameObject.MapCursor
  BattleMap           battle.BattleMap 
  BattleState         SceneState
  BattleFieldBadge    ui.InfoBadge
  CMinionIndex        int
}

func (b *BattleSimulationScene) Update(ticks int) error {
  b.Input()
  b.Cursor.CursorSprite.Xpos = b.Cursor.Xpos
  b.Cursor.CursorSprite.Ypos = b.Cursor.Ypos
  b.BattleMap.Update(ticks)
  return nil
}

func (b *BattleSimulationScene) Draw(screen *ebiten.Image, textures texture.Tex) {
  if b.BattleState == Ready {
    b.BattleMap.Draw(screen, textures)
    b.Cursor.CursorSprite.Draw(screen, textures.Cursor)
    b.BattleFieldBadge.Draw(screen, textures)
  }
}

func (b *BattleSimulationScene) Input() error {
  if ebiten.IsKeyPressed(ebiten.KeyJ) || ebiten.IsKeyPressed(ebiten.KeyS){
    b.Cursor.Ypos += 1 * b.Cursor.Speed
  } else if ebiten.IsKeyPressed(ebiten.KeyK) || ebiten.IsKeyPressed(ebiten.KeyW) {
    b.Cursor.Ypos -= 1 * b.Cursor.Speed
  }
  if ebiten.IsKeyPressed(ebiten.KeyH) || ebiten.IsKeyPressed(ebiten.KeyA) {
    b.Cursor.Xpos -= 1 * b.Cursor.Speed
  } else if ebiten.IsKeyPressed(ebiten.KeyL) || ebiten.IsKeyPressed(ebiten.KeyD) {
    b.Cursor.Xpos += 1 * b.Cursor.Speed
  }
  if inpututil.IsKeyJustPressed(ebiten.KeyI) {
    b.BattleFieldBadge.Msg = b.BattleMap.GetCurrentTileName(
      b.Cursor.Xpos,
      b.Cursor.Ypos,
    )
    b.BattleFieldBadge.Active = !b.BattleFieldBadge.Active
    // log.Printf("BattleFieldBadge Active: %v", b.BattleFieldBadge.Active)
  }
  if inpututil.IsKeyJustPressed(ebiten.KeyX) {
    if index := b.BattleMap.GetTileIndex(b.Cursor.Xpos+10, b.Cursor.Ypos-10); index != -1 {
      log.Printf("Updated DestinationIndex: %d", index)
      for idx := range b.BattleMap.Allies {
        b.BattleMap.Allies[idx].DestinationIndex = index
        b.BattleMap.Allies[idx].TargetIndex = -1
      }
      b.BattleMap.AddEffect(effect.CreateEffect(
        (index%b.BattleMap.Width)*constant.TILESIZE,
        (index/b.BattleMap.Width)*constant.TILESIZE,
        60,
        texture.Animation{
          CurrentAnimationFrame: 0,
          Frames: 6,
        },
        constant.TargetParticle,
      ))
    }
  }
  if ebiten.IsKeyPressed(ebiten.KeyShift) {
    b.Cursor.Speed = 6
  } else {
    b.Cursor.Speed = 3
  }
  return nil  
}

func (b *BattleSimulationScene) Init(screenW, screenH int) error {
  b.sceneW = screenW
  b.sceneH = screenH
  b.State = Starting
  b.BattleState = Starting
  b.Cursor = gameObject.MapCursor{
    Xpos: 0,
    Ypos: 0,
    CursorSprite: gameObject.Sprite{
      Xpos: 0,
      Ypos: 0,
      Width: 16,
      Height: 16,
    },
    Speed: 1,
  }
  b.BattleMap = battle.BattleMap{
    Tiles:    tilemap.StandardTileMap,
    Width:    tilemap.StandardTileMapWidth,
    Effects:  []effect.Effect{},
    Allies:   []minion.Minion{},
    Enemies:  []minion.Minion{},
  }
  // b.BattleMap.Allies = append(b.BattleMap.Allies, minion.InitBaltieMinion(0,0)) 
  b.BattleMap.Allies = append(b.BattleMap.Allies, minion.InitBaltieMinion(0,1)) 
  // b.BattleMap.Allies = append(b.BattleMap.Allies, minion.InitBaltieMinion(0,0)) 
  // b.BattleMap.Allies = append(b.BattleMap.Allies, minion.InitBaltieMinion(0,2)) 
  // b.BattleMap.Enemies = append(b.BattleMap.Enemies, minion.InitRatMinion(9,5)) 
  b.BattleMap.Enemies = append(b.BattleMap.Enemies, minion.InitRatMinion(4,3)) 
  b.BattleMap.Enemies = append(b.BattleMap.Enemies, minion.InitRatMinion(4,4)) 
  b.BattleFieldBadge.Init("FieldInfo", screenW, screenH) 
  b.State = Ready
  b.BattleState = Ready
  return nil
}
