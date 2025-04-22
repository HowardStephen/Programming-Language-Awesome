package main

func main() {
	defer println(recover()) // 无效 nil

	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 有效
		}
	}()

	panic("Panic!")
}
