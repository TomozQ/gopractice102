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
	// myData.Name = "Taro"
	// myData.Data = []int{10, 20, 30}
	// fmt.Println(myData)

	taro := Mydata{
		"Taro",
		[]int{90, 80, 70},
	}
	fmt.Println(taro)
	// taroのポインタを渡す
	rev(&taro)
	fmt.Println(taro)
}

func rev (md *Mydata) {	// ポインタを受ける（Mydata型の値が格納されているポインタ）
	// *mdでポインタの中の値を取得する
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}