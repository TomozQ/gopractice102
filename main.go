package main

import (
	"fmt"
)

/*
	チャンネルの作成
	変数 := make(chan 型)

	チャンネルから値を取り出す
	変数 := <-チャンネル

	値を追加する
	チャンネル <- 値
*/

// total is method.
func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}

func main () {
	// channel生成
	c := make(chan int)
	// 並列処理でtotal関数に100とチャンネルを渡す
	go total(100, c)
	// チャンネルの中身を出力
	fmt.Println("total: ", <-c)
}