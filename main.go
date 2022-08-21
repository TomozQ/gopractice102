package main

import (
	"time"
	"fmt"
	"strconv"
)

func prmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

/*
	c <- s と実行して値を追加
*/
func first(n int, c chan string) {
	const nm string = "first-"
	for i := 0; i < 10; i++ {
		s := nm + strconv.Itoa(i)
		prmsg(n, s)
		c <- s
	}
}

/*
	チャンネルcから取得した値を使ってprmsg関数を実行。
*/
func second(n int, c chan string) {
	for i := 0; i < 10; i ++ {
		prmsg(n, "second:[" + <-c + "]")
	}
}

func main () {
	// チャンネルを作成
	c := make(chan string)
	/*
		firstとsecondに同じチャンネルcを渡す。（複数のスレッド間でチャンネルを利用する場合は、同じチャンネルを使う）
	*/
	go first(10, c)
	second(10, c)
	fmt.Println()
}