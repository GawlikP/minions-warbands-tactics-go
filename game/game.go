package game

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "fmt"
  "log"
  "minions-warbands-tactis/textures"
  "minions-warbands-tactis/scenes"
  "os"
)

type Game struct {
  initialized bool
  tex textures.Tex
  mainMenu scenes.MainMenuScene
  currentScene scenes.CurrentScene
}

func (g *Game) Update() error {
  if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
    os.Exit(1)
  }
  if !g.initialized {
    g.Init()
  }
  if g.currentScene == scenes.MainMenu {
    if g.mainMenu.State != scenes.Ready {
      log.Print("Initializing the MainMenu scene")      
      g.mainMenu.Init(480,320)
      log.Print("Initialized Main Menu scene")      
    }

    if g.currentScene == scenes.MainMenu && g.mainMenu.State == scenes.Ready {
      g.mainMenu.Update()
    }
  }
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  if g.currentScene == scenes.MainMenu && g.mainMenu.State == scenes.Ready {
    g.mainMenu.Draw(screen, g.tex)
  }
  g.PerformanceTab(screen)
}

func (g *Game) Layout(win, hin int) (wout, hout int) {
  return 480, 320
}

func (g *Game) PerformanceTab(screen *ebiten.Image) {
  msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
  ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Init() {
  log.Print("Initializing the game")
  log.Print("Initializing Game Textures")
  g.tex = textures.Tex{}
  g.tex.InitTextures()
  log.Print("GameTexturesInitialized")
  g.mainMenu = scenes.MainMenuScene{
    State: scenes.Closed,
  }
  g.currentScene = scenes.MainMenu
  g.initialized = true
}
