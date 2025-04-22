package main

func main() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()

	panic("Panic!")
}
