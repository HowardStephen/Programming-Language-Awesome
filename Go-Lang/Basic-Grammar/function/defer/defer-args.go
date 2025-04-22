package main

func main() {
  x, y := 10, 20

  println("old: x => ", x, "\ty => ", y)
  
  defer println("defer:\t x => ", x)

  defer func() {
	println("defer closure:\t y => ", y)
  }()

  x += 100
  y += 100

  println("new: x => ", x, "\ty => ", y)
}
