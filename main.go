package main

import (
	"fmt"
	"time"
	"strconv"
)

func main () {
	msg := "start"
	prmsg := func(nm string, n int) {
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	hello := func (n int) {
		const nm string = "hello"
		for i := 0; i < 10; i++ {
			msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}

	main := func (n int) {
		const nm string = "*main"
		for i := 0; i < 5; i++ {
			msg += " m" + strconv.Itoa(i)
			prmsg(nm, n)
		}
	}

	go hello(60)
	main(100)
}

/*
	出力結果
	*main start m0
	hello start m0 h0
	hello start m0 h0 h1
	*main start m0 h0 h1 m1
	hello start m0 h0 h1 m1 h2
	hello start m0 h0 h1 m1 h2 h3
	*main start m0 h0 h1 m1 h2 h3 m2
	hello start m0 h0 h1 m1 h2 h3 m2 h4
	*main start m0 h0 h1 m1 h2 h3 m2 h4 m3
	hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5
	hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6
	*main start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6 m4
	hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6 m4 h7
	hello start m0 h0 h1 m1 h2 h3 m2 h4 m3 h5 h6 m4 h7 h8
*/

/*
	それぞれのスレッドでmsgに値を追加しており、それらは両スレッドでリアルタイムに値を共有している。
	このmsgの値は、それぞれのスレッドで呼び出している関数(helloとmain)が置かれているmain関数ないにあり、これらの関数から利用できるようになっている。
	関数をmain外で定義していると、main内にある変数に直接アクセスすることができない。
*/

/*
	10回繰り返しているのにh9が出力されていない
	→h9が出力される前にmainスレッドが終了している。

	Goルーチンによっていくつものスレッドが実行されるが、それらは全て「メインスレッドが実行中である」間だけ。
	メインスレッドが終了すると、それ以外の全てのスレッドは（処理中であったとしても）全て消えてしまう。
	メインスレッドが終了した後も、残るスレッドが実行し続けるわけではない。
*/