package main

import (
  "log"
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactis/game"
)

func main() {
  ebiten.SetWindowSize(1024, 720)
  ebiten.SetWindowTitle("Minions Warbands Tactics")
  if err := ebiten.RunGame(&game.Game{}); err != nil {
    log.Fatal(err)
  }
}
