package maps 

import (
  "minions-warbands-tactics/gameObjects"
  "log"
)

var StandardTileMap []gameObjects.BattleMapTileType
const StandardTileMapWidth = 10

func InitializeMaps() {
  // blocks := []int{}
  // for i := 0; i < StandardTileMapWidth * 8; i++ {
  //   StandardTileMap = append(StandardTileMap, gameObjects.Grass)
  // }
  //
  // for _, v := range blocks {
  //   StandardTileMap[v] = gameObjects.Stone
  // }

  StandardTileMap = []gameObjects.BattleMapTileType{
    gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
  }

  log.Printf("Standard test map size: %d", len(StandardTileMap))
}
