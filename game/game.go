package game

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "fmt"
)

type Game struct {}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  g.PerformanceTab(screen)
}

func (g *Game) Layout(win, hin int) (wout, hout int) {
  return 480, 320
}

func (g *Game) PerformanceTab(screen *ebiten.Image) {
  msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())
  ebitenutil.DebugPrint(screen, msg)
}


