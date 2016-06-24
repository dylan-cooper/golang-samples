package fib

//Naive recursive definition of fibonacci
func FibFunc(n int) int {
  if n < 2 {
    return 1
  } else {
    return FibFunc(n - 1) + FibFunc(n - 2);
  }
}
