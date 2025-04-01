package main

import "fmt"

type Animal struct {
  name string
}

// Animal's method
func (a *Animal) changeName(name string) {
  a.name = name
}

type Dog struct {
  weight uint8
  Animal
}

// Dog's method
func (d Dog) getInfo() string {
  // 使用 fmt.Sprintf() 将 int 型转换成 string 型
  return d.name + " is " + fmt.Sprintf("%d", d.weight)
}

func main() {
  snoopy := Dog {
    weight: 30,
    Animal: Animal {
      name: "snoopy",
    },
  }

  fmt.Println(snoopy.getInfo())

  snoopy.changeName("...")

  fmt.Println(snoopy.getInfo())
}
