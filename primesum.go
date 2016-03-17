package main

import (
  "fmt"

  "math"
)

func isprime(n int) bool {
  x := 2
  for x < n {
    if math.Mod(float64(n), float64(x)) == 0 {
      return false
    } else {
      x++
    }
  }
  return true
}

func main() {
  sum := 0
  x := 2
  for x <= 1000000 {
    if isprime(x){
      sum += x
    }
    x++
  }
  fmt.Println(sum)
}
