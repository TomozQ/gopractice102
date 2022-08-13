package main

import (
	"fmt"
)

// どんな値でも保管できる型
type General interface {}

func main () {
	var v General
	v = 123
	fmt.Println(v)
	v = 0.01
	fmt.Println(v)
	v = "Hello"
	fmt.Println(v)
	v = true
	fmt.Println(v)
}