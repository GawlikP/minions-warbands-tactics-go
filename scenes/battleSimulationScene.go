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
    b.Cursor.Ypos += 1
  } else if ebiten.IsKeyPressed(ebiten.KeyK) {
    b.Cursor.Ypos -= 1
  }
  if ebiten.IsKeyPressed(ebiten.KeyH) {
    b.Cursor.Xpos -= 1
  } else if ebiten.IsKeyPressed(ebiten.KeyL) {
    b.Cursor.Xpos += 1
  }
  if inpututil.IsKeyJustPressed(ebiten.KeyI) {
    b.BattleFieldBadge.Msg = b.BattleMap.GetCurrentTileName(
      b.Cursor.Xpos,
      b.Cursor.Ypos,
    )
    b.BattleFieldBadge.Active = !b.BattleFieldBadge.Active
    log.Printf("BattleFieldBadge Active: %v", b.BattleFieldBadge.Active)
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
  }
  b.BattleMap = gameObjects.BattleMap{
    Units: []gameObjects.Unit{{},{}},
    Tiles: maps.StandardTileMap,
    Width: maps.StandardTileMapWidth,
  }
  b.BattleFieldBadge.Init("FieldInfo", screenW, screenH) 
  b.State = Ready
  b.BattleState = Ready
  return nil
}
