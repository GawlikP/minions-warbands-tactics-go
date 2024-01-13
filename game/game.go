package game

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/inpututil"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "fmt"
  "log"
  "minions-warbands-tactics/textures"
  "minions-warbands-tactics/scenes"
  "minions-warbands-tactics/maps"
  "os"
)

type Game struct {
  initialized       bool
  tex               textures.Tex
  mainMenu          scenes.MainMenuScene
  battleSimulation  scenes.BattleSimulationScene
  currentScene      scenes.CurrentScene
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
  if g.currentScene == scenes.MainMenu && g.mainMenu.State == scenes.Ready {
    g.mainMenu.Draw(screen, g.tex)
  }
  if g.currentScene == scenes.BattleSimulation && g.battleSimulation.State == scenes.Ready {
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
  if g.currentScene == scenes.MainMenu {
    if g.mainMenu.State != scenes.Ready {
      log.Print("Initializing the MainMenu scene")      
      g.mainMenu.Init(480,320)
      log.Print("Initialized Main Menu scene")      
    }

    if g.mainMenu.State == scenes.Ready {
      menuOutput := g.mainMenu.Update()

      if menuOutput == 1 {
        log.Print("Changing scene to BattleSimulation")
        g.mainMenu.State = scenes.Closed
        g.currentScene = scenes.BattleSimulation
      }
    }
  }
}


func (g *Game) ProcessBattleSimulator() {
  if g.currentScene == scenes.BattleSimulation {
    if g.battleSimulation.State != scenes.Ready {
      log.Print("Initializing the BattleSimulation scene")      
      g.battleSimulation.Init(480,320)
      log.Print("Initialized BattleSimulation scene")      
    }

    if g.battleSimulation.State == scenes.Ready {
      g.battleSimulation.Update(g.ticks)
    }
  }
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
  log.Print("Initializing MAPS")
  maps.InitializeMaps()
  g.battleSimulation = scenes.BattleSimulationScene{
    State: scenes.Closed,
    BattleState: scenes.Closed,
  }
  g.currentScene = scenes.MainMenu
  g.initialized = true
}

func (g *Game) UpdateTicks() {
  g.ticks += 1
  g.ticks %= 60
}
