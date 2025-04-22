package main

import "fmt"

func generateFibonacci(number int) int {
  if number == 1 || number == 0 {
    return number
  }
  return generateFibonacci(number - 1) + generateFibonacci(number - 2)
}

func main() {
  for number := 0; number < 10; number++ {
    fmt.Printf("%d ", generateFibonacci(number))
  }
  fmt.Println()
}
