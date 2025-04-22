package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int64 = 10

	t := reflect.TypeOf(a)
	v := reflect.ValueOf(&a)

	fmt.Println(t, t.Name(), t.Kind())
	fmt.Println(v, v.Kind(), v.Interface())

	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(100)
	}
	fmt.Println(a)
}
