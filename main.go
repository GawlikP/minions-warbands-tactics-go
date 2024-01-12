package main

import (
  "log"
  "github.com/hajimehoshi/ebiten/v2"
  "minions-warbands-tactics/game"
)

func main() {
  ebiten.SetWindowSize(1980, 1080)
  ebiten.SetWindowTitle("Minions Warbands Tactics")
  if err := ebiten.RunGame(&game.Game{}); err != nil {
    log.Fatal(err)
  }
}
