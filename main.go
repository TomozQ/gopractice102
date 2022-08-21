package main

import (
	"fmt"
)

func total (cs chan int, cr chan int) {
	n := <-cs
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	cr <- t
}

func main () {
	// チャンネルを二つ作成
	cs := make(chan int)
	cr := make(chan int)
	// トータル関数にチャンネルを二つ渡す
	go total(cs, cr)
	cs <- 100
	fmt.Println("total:", <- cr)
}

/*
	出力結果
	n =  100
	total: 5050
*/

/*
	total関数ではチャンネルからcsの値を取り出し、演算の結果をcrに追加
	main関数ではチャンネルcsに値を追加し、チャンネルcrから取り出した値を表示
	...csをmain→total
	...crをtotal→main
	というようにして、二つのチャンネルで双方向に値を送受していた。
*/