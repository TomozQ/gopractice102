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

/*
	出力結果
	total:  5050
*/

/*
	go total
	fmt.Println()

	とtotal関数を並列処理で実行しているのに
	常にfmt.Printlnが問題なく実行されている。

	通常はfmt.Printlnに処理が入り、このメインスレッドの処理が終了すればtotal関数を実行しているスレッドも中断される。
	が、実際はそうはならない。
	「チャンネルから値を取得する場合、その値が送られてくるまで処理を待つ」
*/