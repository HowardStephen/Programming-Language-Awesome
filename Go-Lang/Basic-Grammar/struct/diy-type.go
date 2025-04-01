package main

import "fmt"

// 自定义一个 uint8 类型的类型
type age uint8

func main() {
  // 声明一个 age 类型的变量
  var henryAge age = 20

  fmt.Printf("type: %T, value: %v\n", henryAge, henryAge)
}
