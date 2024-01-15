package texture

import (
  "log"
)

func InitMinionsTextures(tex *Tex) { 
  log.Print("Initializing Bazaltie Textures")
  tex.BazaltieWalkingRight = LoadAnimationFrames("Bazaltie", "Side", 4, false)
  tex.BazaltieWalkingLeft = LoadAnimationFrames("Bazaltie", "Side", 4, true)
  tex.BazaltieWalkingUp = LoadAnimationFrames("Bazaltie", "Up", 4, false)
  tex.BazaltieWalkingDown = LoadAnimationFrames("Bazaltie", "Down", 4, false)
  log.Print("Initializing Threedy Textures")
  tex.ThreedyWalkingRight = LoadAnimationFrames("Threedy", "Side", 4, false)
  tex.ThreedyWalkingLeft = LoadAnimationFrames("Threedy", "Side", 4, true)
  tex.ThreedyWalkingUp = LoadAnimationFrames("Threedy", "Up", 4, false)
  tex.ThreedyWalkingDown = LoadAnimationFrames("Threedy", "Down", 4, false)
  log.Print("Initialized the Minions Textures")
}
