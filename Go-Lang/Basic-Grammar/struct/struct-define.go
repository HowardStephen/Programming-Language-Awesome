package main

import "fmt"

type Person struct {
  // 小写字母开头 - 内部使用
  age uint8
  name string

  // 大写字母开头 - 外部使用
  Id uint
  Address string
}

func main() {
  user := Person {
    age: 22,
    name: "henry",
    Id: 20,
    Address: "new york",
  }

  fmt.Printf("age: %v,\nname: %v,\nId: %v,\nAddress: %v\n", user.age, user.name, user.Id, user.Address)
}
