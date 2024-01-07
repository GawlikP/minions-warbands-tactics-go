package scenes

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "minions-warbands-tactics/gameObjects"
  "minions-warbands-tactics/textures"
)

type MainMenuScene struct {
  CursorPosition      int
  MaxCursorPosition   int
  CursorSprite        gameObjects.Sprite
  State               SceneState
  sceneW              int
  sceneH              int
  NewGameSprite       gameObjects.Sprite
  ExitGameSprite      gameObjects.Sprite
  NewGameChoosed      bool
}

func (m *MainMenuScene) Update() int {
  m.Input()
  if m.CursorPosition > m.MaxCursorPosition {
    m.CursorPosition = 1
  }
  if m.CursorPosition < 0 {
    m.CursorPosition = 0
  }

  if m.State == Ready { 
    if m.CursorPosition == 1 {
      m.CursorSprite.Ypos =  m.NewGameSprite.Ypos + 32
    } else {
      m.CursorSprite.Ypos = m.ExitGameSprite.Ypos + 32
    }
  }

  if m.NewGameChoosed {
    return 1
  }
  return 0
}

func (m *MainMenuScene) Draw(screen *ebiten.Image, textures textures.Tex) {
  if m.State == Ready { 
    m.NewGameSprite.Draw(screen, textures.NewGameBanner)
    m.ExitGameSprite.Draw(screen, textures.ExitGameBanner)
    m.CursorSprite.Draw(screen, textures.Cursor)
  }
}

func (m *MainMenuScene) Input() error {
  if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
    m.CursorPosition += 1
  } else if inpututil.IsKeyJustPressed(ebiten.KeyK) {
    m.CursorPosition -= 1
  }
  if inpututil.IsKeyJustPressed(ebiten.KeySpace) && m.CursorPosition == 1 { 
    m.NewGameChoosed = true
  }
  return nil
}

func (m *MainMenuScene) Init(screenW, screenH int) {
  m.sceneW = screenW
  m.sceneH = screenH
  m.State = Starting
  m.CursorPosition = 0
  m.MaxCursorPosition = 1
  m.CursorSprite = gameObjects.Sprite{
    Width: 16,
    Height: 16,
    Xpos: screenW/2,
    Ypos: screenH/2,
  }
  m.NewGameSprite = gameObjects.Sprite{
    Width: 240,
    Height: 64,
    Xpos: screenW/2 - 120,
    Ypos: screenH/2-64,
  }
  m.ExitGameSprite = gameObjects.Sprite{
    Width: 240,
    Height: 64,
    Xpos: screenW/2 - 120,
    Ypos: screenH/2+64,
  }
  m.NewGameChoosed = false
  m.State = Ready
}
