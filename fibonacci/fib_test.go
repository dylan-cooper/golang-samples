package fib

import "testing"

var tests = []int{1,1,2,3,5,8,13,21,34,55,89}

func TestFibFunc(t *testing.T) {

  for i, v := range tests {
    if val := FibFunc(i); v != val {
      t.Fatalf("at index %d, expected %d got %d", i, v, val);
    }
  }
}
