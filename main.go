package main

import (
	"fmt"
)

// 構造体 複数の値をひとまとめにした値
// struct {
// 	'''変数宣言'''

// }

// var 変数 struct { ...... }

// var myData struct {
// 	Name string
// 	Data []int
// }

// 構造体をtypeで型として定義する
// type 型名 struct {'''内容'''}
// →　作成... 型名 {'''値の設定'''}

// Mydata is structure 			「外部から利用可能」なもの（大文字で始まる名前）はコメントが必須
type Mydata struct {
	Name string
	Data []int
}

func main () {
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 5, 5)
	fmt.Println(taro)
}