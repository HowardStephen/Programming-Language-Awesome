package main

import "fmt"

func numberCalc(number int) (func(int) int, func(int)int) {
  add := func(num int) int {
    number += num
    return number
  }

  sub := func(num int) int {
    number -= num
    return number
  }

  return add, sub
}

func main() {
  add, sub := numberCalc(10)

  fmt.Println(add(5), sub(10))

  fmt.Println(add(10), sub(5))
}
