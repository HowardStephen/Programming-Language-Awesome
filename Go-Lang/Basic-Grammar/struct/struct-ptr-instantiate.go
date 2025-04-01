package main

import "fmt"

type User struct {
  name string
  id uint
}

// new() 写法 - 创建指针结构体
func newCreateUser(name string, id uint) *User {
  user := new(User)
  user.name = name
  user.id = id

  return user
}

// 取地址写法 - 创建指针结构体
func ptrCreateUser(name string, id uint) *User {

  user := &User {
    name,
    id, // 不要漏掉最后一个,
  }
  return user
}

func main() {
  user1 := newCreateUser("henry", 20)
  user2 := ptrCreateUser("green", 20)
  fmt.Printf("name: %v, id: %v\n", user1.name, user1.id)
  fmt.Printf("name: %v, id: %v\n", user2.name, user2.id)
}
