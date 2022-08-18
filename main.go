package main

import (
	"fmt"
)

/*
	チャンネルには複数の値を保管することができるが
	これは先入先出となっている。
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
	go total(1000, c)
	go total(100, c)
	go total(10, c)

	x,y,z := <-c, <-c, <-c
	// チャンネルの中身を出力
	fmt.Println(x, y, z)
}

/*
	出力結果
	55 500500 5050
*/

/*
	先入先出なので「実行したスレッドにかかる時間」によって
	実行時間が短いものから順に c <- t が実行され、スレッドが終了する。
	「必ずGoルーチンを呼び出した順に値が追加されるわけではない」
*/