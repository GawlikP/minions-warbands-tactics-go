package scenes

type SceneState int

const (
  Closed SceneState = iota
  Ready
  Starting
)

type CurrentScene int

const (
  MainMenu CurrentScene = iota
  BattleSimulation // to do in 0.1
)
