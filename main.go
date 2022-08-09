package main

import (
	"fmt"
)

// 構造体 複数の値をひとまとめにした値
// struct {
// 	'''変数宣言'''

// }

// var 変数 struct { ...... }

var myData struct {
	Name string
	Data []int
}

func main () {
	myData.Name = "Taro"
	myData.Data = []int{10, 20, 30}
	fmt.Println(myData)
}