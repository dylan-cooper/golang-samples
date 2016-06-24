package vector

import ("math")

type Vector struct {
  X, Y float64
}

func (v Vector) Magnitude () float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
