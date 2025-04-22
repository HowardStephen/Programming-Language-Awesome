package main

import "fmt"

type echoMan struct {
	content string
}

func (e echoMan) echo(name string) {
	fmt.Println(e.content, name)
}

func main() {
	man := echoMan{
		content: "Hello",
	}
	echoMan.echo(man, "world")
	
}
