package main

import "fmt"

// 类型别名
type age = uint8

func main() {
  // 声明一个 age 别名的变量
  var henryAge age = 20

  fmt.Printf("type: %T, value: %v\n", henryAge, henryAge)
}

