package tilemap

import (
  "minions-warbands-tactics/constant"
  "log"
)

var StandardTileMap []constant.BattleMapTileType
const StandardTileMapWidth = 10

func InitializeStandardMap() {
  // blocks := []int{}
  // for i := 0; i < StandardTileMapWidth * 8; i++ {
  //   StandardTileMap = append(StandardTileMap, constant.Grass)
  // }
  //
  // for _, v := range blocks {
  //   StandardTileMap[v] = constant.Stone
  // }

  StandardTileMap = []constant.BattleMapTileType{
    constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass,
    constant.Grass, constant.Stone, constant.Stone, constant.Stone, constant.Stone, constant.Stone, constant.Stone, constant.Grass, constant.Grass, constant.Grass,
    constant.Grass, constant.Grass, constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Grass,
    constant.Grass, constant.Stone, constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Grass,
    constant.Grass, constant.Stone, constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Grass,
    constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Grass,
    constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Grass, constant.Grass, constant.Stone, constant.Grass, constant.Grass, constant.Grass,
    constant.Grass, constant.Stone, constant.Stone, constant.Stone, constant.Stone, constant.Stone, constant.Stone, constant.Grass, constant.Grass, constant.Grass,
  }

  log.Printf("Standard test map size: %d", len(StandardTileMap))
}
