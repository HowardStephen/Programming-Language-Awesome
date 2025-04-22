package main

func test1(num int) {
  println("func test1 ==>")
  defer println("a")
  defer println("b")
  defer println(100 / num)
  defer println("c")
}

func test2(num int) {
  println("func test2 ==>")
  defer println("a")
  defer println("b")
  defer func() {
    println(100 / num)
  }()
  defer println("c")
}

func main() {
  // test1(0)
  
  test2(0)
}
