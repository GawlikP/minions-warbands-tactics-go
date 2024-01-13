package maps 

import (
  "minions-warbands-tactics/gameObject"
  "log"
)

var StandardTileMap []gameObject.BattleMapTileType
const StandardTileMapWidth = 10

func InitializeMaps() {
  // blocks := []int{}
  // for i := 0; i < StandardTileMapWidth * 8; i++ {
  //   StandardTileMap = append(StandardTileMap, gameObject.Grass)
  // }
  //
  // for _, v := range blocks {
  //   StandardTileMap[v] = gameObject.Stone
  // }

  StandardTileMap = []gameObject.BattleMapTileType{
    gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass,
    gameObject.Grass, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass,
    gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass,
    gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass,
    gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass,
    gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass,
    gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Grass, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass,
    gameObject.Grass, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Stone, gameObject.Grass, gameObject.Grass, gameObject.Grass,
  }

  log.Printf("Standard test map size: %d", len(StandardTileMap))
}
