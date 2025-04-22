package main

import "fmt"

/*
* slice 和 array 的区别:
* 	  slice: []Type
* 	  array: [size]Type
 */

func main() {
	s1 := []int{1, 2, 3}
	fmt.Printf("s type: %T\n", s1)

	s2 := [...]int{1, 2, 3}
	fmt.Printf("s type: %T\n", s2)
}
