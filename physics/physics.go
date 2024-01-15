package physics

import "math"

func IsInLine(x1, y1, x2, y2, x3, y3 int) bool {
  return (x1 - x3)*(y3 - y2) == (x3 - x2)*(y1 - y3)
}

func IsColidingOnCircle(x, y, w, h, cx, cy, r int) bool {
  closeX := -1
  closeY := -1
  if cx < x {
    closeX = x
  }
  if cx > x  + w {
    closeX = x + w
  }
  if cy < y {
    closeY =  y
  }
  if cy > y + h {
    closeY = y + h
  }

  dis := math.Sqrt(float64((closeX - cx)*(closeX - cx) + (closeY - cy)*(closeY - cy)))

 if dis < float64(r) {
   return true
 }
 return false
}
