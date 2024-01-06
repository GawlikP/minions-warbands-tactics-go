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
  WarSimulation // to do in 0.1
)
