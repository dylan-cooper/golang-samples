package lcm

func LCM(a int, b int) int {
  r := a * b
  //euclidean algorithm for gcd
  for b != 0 {
    a, b = b, a % b
  }
  //divide product by gcd to get lcm
  return r / a
}
