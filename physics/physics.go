package physics

import "math"

func IsInLine(x1, y1, x2, y2, x3, y3 float64) bool {
  //  dist_line_endp = DistBetwPoints(A,B)
  // if DistBetwPoints(A,C)>dist_line_endp:       return 1
  // elif DistBetwPoints(B,C)>dist_line_endp:     return 1
  // else:                                        return 0
  dis := Dis(x1,y1,x2,y2)

  if Dis(x1, y1, x3, y3 ) > dis {
    return false
  }
  if Dis(x2,y2, x3, y3) > dis {
    return false
  }
  return true
}

func IsColidingOnCircle(x, y, w, h, cx, cy, r int) bool {
  closeX := cx
  closeY := cy
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

func Dis(x1, y1, x2, y2 float64) float64 {
  return math.Sqrt(math.Pow((x2 - x1),2) + math.Pow((y2 - x2),2))
}
