package texture

import (
  "log"
)

func InitParticlesTextres(tex *Tex) { 
  log.Printf("Initializing Standard Particle Textures")
  tex.StandardParticle = LoadAnimationFrames("Particles","Standard", 6, false)
  log.Printf("Initializing Target Particle Textures")
  tex.TargetParticle = LoadAnimationFrames("Particles","Target", 6, false)
}
