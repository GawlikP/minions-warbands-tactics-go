package maps 

import (
  "minions-warbands-tactics/gameObjects"
  "log"
)

var StandardTileMap []gameObjects.BattleMapTileType
const StandardTileMapWidth = 8

func InitializeMaps() {
  // for i := 0; i < StandardTileMapWidth * 8; i++ {
  //   StandardTileMap = append(StandardTileMap, gameObjects.Grass)
  // }

  StandardTileMap = []gameObjects.BattleMapTileType{
    gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Sand,  gameObjects.Sand,  gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Sand,  gameObjects.Sand,  gameObjects.Grass, gameObjects.Grass, 
    gameObjects.Grass, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Sand,  gameObjects.Grass,
    gameObjects.Grass, gameObjects.Stone, gameObjects.Stone, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Sand,  gameObjects.Grass, 
    gameObjects.Grass, gameObjects.Stone, gameObjects.Stone, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
    gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass, gameObjects.Grass,
  }

  log.Printf("Standard test map size: %d", len(StandardTileMap))
}
