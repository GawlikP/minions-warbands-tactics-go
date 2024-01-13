package game

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "fmt"
  "log"
  "minions-warbands-tactics/texture"
  "minions-warbands-tactics/scene"
  "minions-warbands-tactics/maps"
  "os"
)

type Game struct {
  initialized       bool
  tex               texture.Tex
  mainMenu          scene.MainMenuScene
  battleSimulation  scene.BattleSimulationScene
  currentScene      scene.CurrentScene
  ticks             int
}

func (g *Game) Update() error {
  g.UpdateTicks()
  if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
    os.Exit(1)
  }
  if !g.initialized {
    g.Init()
  }

  g.ProcessMainMenu()
  g.ProcessBattleSimulator()

  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  if g.currentScene == scene.MainMenu && g.mainMenu.State == scene.Ready {
    g.mainMenu.Draw(screen, g.tex)
  }
  if g.currentScene == scene.BattleSimulation && g.battleSimulation.State == scene.Ready {
    g.battleSimulation.Draw(screen, g.tex)
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

func (g *Game) ProcessMainMenu() {
  if g.currentScene == scene.MainMenu {
    if g.mainMenu.State != scene.Ready {
      log.Print("Initializing the MainMenu scene")      
      g.mainMenu.Init(480,320)
      log.Print("Initialized Main Menu scene")      
    }

    if g.mainMenu.State == scene.Ready {
      menuOutput := g.mainMenu.Update()

      if menuOutput == 1 {
        log.Print("Changing scene to BattleSimulation")
        g.mainMenu.State = scene.Closed
        g.currentScene = scene.BattleSimulation
      }
    }
  }
}


func (g *Game) ProcessBattleSimulator() {
  if g.currentScene == scene.BattleSimulation {
    if g.battleSimulation.State != scene.Ready {
      log.Print("Initializing the BattleSimulation scene")      
      g.battleSimulation.Init(480,320)
      log.Print("Initialized BattleSimulation scene")      
    }

    if g.battleSimulation.State == scene.Ready {
      g.battleSimulation.Update(g.ticks)
    }
  }
}

func (g *Game) Init() {
  log.Print("Initializing the game")
  log.Print("Initializing Game texture")
  g.tex = texture.Tex{}
  g.tex.InitTextures()
  log.Print("GametextureInitialized")
  g.mainMenu = scene.MainMenuScene{
    State: scene.Closed,
  }
  log.Print("Initializing MAPS")
  maps.InitializeMaps()
  g.battleSimulation = scene.BattleSimulationScene{
    State: scene.Closed,
    BattleState: scene.Closed,
  }
  g.currentScene = scene.MainMenu
  g.initialized = true
}

func (g *Game) UpdateTicks() {
  g.ticks += 1
  g.ticks %= 60
}
