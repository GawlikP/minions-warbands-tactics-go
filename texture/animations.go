package texture

type Animation struct {
  CurrentAnimationFrame     int
  Frames                    int
}

func (a *Animation) UpdateAnimationIndex(tick int, animationLength int) {
  if a.Frames == 0 {
      return
  }
  index := tick/(60/a.Frames)

  if index < animationLength {
    a.CurrentAnimationFrame = index
  }
}

func InitAnimation() Animation {
  a := Animation{}
  a.CurrentAnimationFrame = 0
  a.Frames = 0 
  return a
}


