package main

import "fmt"

type User struct {
  string // 匿名字段 （唯一）
  uint8
}

func main() {
  // 匿名结构体
  var user1 struct{ name string; age int }
  user1.name = "henry"
  user1.age = 20

  var user2 = User {
    "green",
    22,
  }

  fmt.Printf("{ name: %v, age: %v }\n", user1.name, user1.age)

  // 匿名字段默认采用类型名作为字段名
  fmt.Printf("{ name: %v, age: %v }\n", user2.string, user2.uint8)
}
