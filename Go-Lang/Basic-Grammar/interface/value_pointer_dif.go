package main

type Actor interface {
	bark()
	move()
}

type Dog struct {
	name  string
	speed int
	voice string
}

func (d Dog) bark() {
	println("It will", d.voice)
}

func (d Dog) move() {
	println("It walks", d.speed)
}

type Cat struct {
	name  string
	speed int
	voice string
}

func (c *Cat) bark() {
	println("It will", c.voice)
}

func (c *Cat) move() {
	println("It walks", c.speed)
}

func main() {
	// region Dog struct
	dog := Dog{
		name:  "snoopy",
		speed: 10,
		voice: "汪汪",
	}
	dog.move()
	dog.bark()
	// endregion

	// region Cat struct
	cat := Cat{
		name:  "tom",
		speed: 20,
		voice: "喵喵",
	}
	cat.move()
	cat.bark()
	// endregion

	var dogActor Actor = dog
	var catActor Actor = &cat // can not use cat, but &cat

	dogActor.move()
	dogActor.bark()
	catActor.move()
	catActor.bark()

}
