package scenes

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "minions-warbands-tactics/gameObjects"
  "minions-warbands-tactics/textures"
  "minions-warbands-tactics/maps"
  "minions-warbands-tactics/ui"
  "log"
)

type BattleSimulationScene struct {
  State               SceneState
  sceneW              int
  sceneH              int
  Cursor              gameObjects.MapCursor
  BattleMap           gameObjects.BattleMap 
  BattleState         SceneState
  BattleFieldBadge    ui.InfoBadge
}

func (b *BattleSimulationScene) Update() error {
  b.Input()
  b.Cursor.CursorSprite.Xpos = b.Cursor.Xpos
  b.Cursor.CursorSprite.Ypos = b.Cursor.Ypos
  b.BattleMap.Update()
  return nil
}

func (b *BattleSimulationScene) Draw(screen *ebiten.Image, textures textures.Tex) {
  if b.BattleState == Ready {
    b.BattleMap.Draw(screen, textures)
    b.Cursor.CursorSprite.Draw(screen, textures.Cursor)
    b.BattleFieldBadge.Draw(screen, textures)
  }
}

func (b *BattleSimulationScene) Input() error {
  if ebiten.IsKeyPressed(ebiten.KeyJ) {
    b.Cursor.Ypos += 1 * b.Cursor.Speed
  } else if ebiten.IsKeyPressed(ebiten.KeyK) {
    b.Cursor.Ypos -= 1 * b.Cursor.Speed
  }
  if ebiten.IsKeyPressed(ebiten.KeyH) {
    b.Cursor.Xpos -= 1 * b.Cursor.Speed
  } else if ebiten.IsKeyPressed(ebiten.KeyL) {
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
    if index := b.BattleMap.GetTileIndex(b.Cursor.Xpos, b.Cursor.Ypos); index != -1 {
      log.Printf("Updated TargetIndex: %d", index)
      b.BattleMap.Minions[0].TargetIndex = index
    }
  }
  if ebiten.IsKeyPressed(ebiten.KeyShift) {
    b.Cursor.Speed = 10
  } else {
    b.Cursor.Speed = 1
  }
  return nil  
}

func (b *BattleSimulationScene) Init(screenW, screenH int) error {
  b.sceneW = screenW
  b.sceneH = screenH
  b.State = Starting
  b.BattleState = Starting
  b.Cursor = gameObjects.MapCursor{
    Xpos: 0,
    Ypos: 0,
    CursorSprite: gameObjects.Sprite{
      Xpos: 0,
      Ypos: 0,
      Width: 16,
      Height: 16,
    },
    Speed: 1,
  }
  b.BattleMap = gameObjects.BattleMap{
    Minions: []gameObjects.Minion{},
    Tiles: maps.StandardTileMap,
    Width: maps.StandardTileMapWidth,
  }
  b.BattleMap.Minions = append(b.BattleMap.Minions, gameObjects.InitRatMinion(0,0)) 
  b.BattleMap.Minions = append(b.BattleMap.Minions, gameObjects.InitFishMinion(6,4)) 
  b.BattleFieldBadge.Init("FieldInfo", screenW, screenH) 
  b.State = Ready
  b.BattleState = Ready
  return nil
}
