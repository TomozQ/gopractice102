package main

import (
	"time"
	"fmt"
)

/*
	チャンネルには複数の値を保管することができるが
	これは先入先出となっている。
*/

// total is method.
func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

func main () {
	// channel生成
	c := make(chan int)
	// c <- 100
	// go total(c)
	
	// Goルーチンでtotal関数を実行
	go total(c)
	/*
		total関数を実行してからチャンネルに100を代入しているが
		チャンネルはまだ値が用意されていない場合は、送られてくるまで処理を待つので
		チャンネルに100が入ってからtotal関数の n := <-c が実行される。
	*/

	c <- 100
	time.Sleep(100 * time.Millisecond)
}

/*
	出力結果
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send]:
	main.main()

	exit status 2
*/

/*
	チャンネルは双方で準備できてから使う
	チャンネルは複数のスレッド間で値をやり取りするためのもの
	これを正常に動作させるには値を送る側と受け取る側の双方向で値の準備ができていなければならない。
	つまり、Goルーチンを実行した後でないとチャンネルは使えない。
*/