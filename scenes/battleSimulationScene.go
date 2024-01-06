package scenes

import (
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/gameObjects"
  "minions-warbands-tactics/textures"
  "minions-warbands-tactics/maps"
)
type BattleSimulationScene struct {
  State               SceneState
  sceneW              int
  sceneH              int
  Cursor              gameObjects.MapCursor
  BattleMap           gameObjects.BattleMap 
}

func (b *BattleSimulationScene) Update() error {
  return nil
}

func (b *BattleSimulationScene) Draw(screen *ebiten.Image, textures textures.Tex) {}

func (b *BattleSimulationScene) Input() error {
  return nil  
}

func (b *BattleSimulationScene) Init(screenW, screenH int) error {
  b.sceneW = screenW
  b.sceneH = screenH
  b.State = Starting
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
  b.State = Ready
  return nil
}
