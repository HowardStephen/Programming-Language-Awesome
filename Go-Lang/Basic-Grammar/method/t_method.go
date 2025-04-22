package main

import "fmt"

type T struct{}

func (t T) ValueMethod() {
	fmt.Println("ValueMethod called")
}

func (t *T) PointerMethod() {
	fmt.Println("PointerMethod called")
}

func main() {
	var t T
	var pt *T = &t

	t.ValueMethod()   // ✅ 可以，T 方法集包含 ValueMethod
	t.PointerMethod() // ❌ 报错，T 方法集不包含 PointerMethod - 没有报错

	pt.ValueMethod()   // ✅ 可以，自动解引用 (*pt).ValueMethod()
	pt.PointerMethod() // ✅ 可以，*T 方法集包含 PointerMethod
}
