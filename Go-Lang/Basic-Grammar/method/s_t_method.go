package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) echoInfo() {
	fmt.Println("I am", a.name)
}

func (a *Animal) changeName(name string) {
	a.name = name
}

type Cat struct {
	Animal `json:"animal"`
}

func (c Cat) echotype() {
	fmt.Println("I am cat")
}

func main() {
	cat := Cat{Animal{"cat"}}
	cat.echoInfo()
	cat.echotype()
	cat.changeName("catppuccin")
	cat.echoInfo()
}
