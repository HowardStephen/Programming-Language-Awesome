package main

import "fmt"

type User struct {
  id uint8
  name string
  vip bool
}

// method
// 值类型的接收者
func (u User) printUser() {
  fmt.Println("todo)) ==>")
  fmt.Printf("id:\t%v\nname:\t%v\nvip:\t%v\n\n", u.id, u.name, u.vip)
}

// 指针类型的接受者
func (u *User) beVip() {
  if u.vip == false {
    u.vip = !u.vip
  }
}

func main() {
  user := User {
    id: 20,
    name: "henry",
    vip: true,
  }

  user.printUser()

  user1 := User {
    id: 19,
    name: "green",
    vip: false,
  }

  user1.printUser()
  user1.beVip()
  user1.printUser()
}
