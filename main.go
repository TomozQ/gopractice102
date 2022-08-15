package main

import (
	"fmt"
	"time"
)

func hello(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>", i, s)
		// nミリ秒だけ処理を停止させる。
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func main () {
	go hello("hello", 50)
	hello("bye!", 100)
}

/*
	出力結果
	<1 bye!><1 hello><2 hello><2 bye!><3 hello><4 hello><3 bye!><5 hello><6 hello><4 bye!>...

	helloを50ミリ秒間隔で実行
	byeを100ミリ秒間隔で実行
*/

/*
	並行処理と非同期処理
	「Goは文法レベルで並行処理を実現している」というのは、Goの大きな特徴の一つ
	「そんなの他の言語でも普通にできるんじゃない？」と思うかもしれない。
	例えばWebなどで多くの人に馴染みのあるJavaScriptはタイマー機能やAjaxといった機能で同時に複数の処理が行えるように見える。
	が、実はこれは「並行処理」ではなく「非同期処理」。
	同期せずに処理を実行できるということであり、実際に実行している処理は常に一つ。
	順番に処理を実行しているところに「これをやって」と処理を割り込ませているだけ。
	Goルーチンは、メインで実行される処理（スレッド）とは別に、独立したスレッドで処理を実行する。
	非同期処理とは根本的に異なるもの
*/