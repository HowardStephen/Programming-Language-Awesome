package main

func main() {
	defer func() {
		println(recover().(string))
	}()

	defer func() {
		panic("Second Panic!")
	}()

	panic("First Panic!")
}
