package lcm

import "testing"

var lcm_tests = []struct {
  a, b, expected int
}{
  {5, 7, 35},
  {8, 9, 72},
  {2, 8, 8},
}

func TestLCM(t *testing.T) {
  for _, test := range lcm_tests {
    if v:= LCM(test.a, test.b); v != test.expected {
      t.Errorf("LCM(%d, %d) returned %d, expected %d", test.a, test.b, v, test.expected)
    }
  }
}
