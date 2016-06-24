package vector

import "testing"
import "math"

var vector_tests = []struct {
  x, y, expected float64
}{
  {1, 1, math.Sqrt(2)},
  {3, 4, 5},
  {12, 5, 13},
}

func TestVectorMagnitude (t *testing.T) {
  for _, test := range vector_tests {
    if v := (Vector{test.x, test.y}).Magnitude(); v != test.expected {
      t.Errorf("{%f, %f}.Magnitude() returned %f, expected %f", test.x, test.y, v, test.expected)
    }
  }
}
