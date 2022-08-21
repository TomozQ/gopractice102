package main

import (
	"time"
	"fmt"
)

/*
	select {
	case 文:
		'''実行する処理'''
	case 文:
		'''実行する処理'''
	
	'''必要なだけcaseを用意'''

	default:
		'''全て当てはまらない場合の処理'''
	}

	caseにはチャンネル操作の文を用意する
	基本的には「←チャンネル」という形でチャンネルから値を取り出す文を記述しておく。

	selectに進むとチャンネルから値を取り出せたcaseに進む。
	複数のチャンネルで値が取り出せる場合には、その中からランダムにcaseが選ばれる。
	全てのチャンネルで準備が整わない場合には default文 が実行される
*/

func count(n int, s int, c chan int) {
	for i := 1; i <= n; i++ {
		c <- i
		time.Sleep(time.Duration(s) * time.Millisecond)
	}
}

/*
	三つのスレッドでcount関数を実行
*/
func main() {
	n1, n2, n3 := 3, 5, 10
	m1, m2, m3 := 100, 75, 50
	c1 := make(chan int)
	go count(n1, m1, c1)
	c2 := make(chan int)
	go count(n2, m2, c2)
	c3 := make(chan int)
	go count(n3, m3, c3)

	// 最後まで値が取り出せるようにn1,n2,n3を全て足した数分ループを回す
	for i := 0; i < n1 + n2 + n3; i++ {
		// チャンネルに値を送るとそのチャンネルのcaseにジャンプし、送られた値を変数reに取り出す。
		select {
		case re := <-c1:
			fmt.Println("*   first  ", re)
		case re := <-c2:
			fmt.Println("**  second ", re)
		case re := <-c3:
			fmt.Println("*** third  ", re)
		}
	}

	fmt.Println("*** finish ***")
}

/*
	出力結果
	*** third   1
	**  second  1
	*   first   1
	*** third   2
	**  second  2
	*   first   2
	*** third   3
	*** third   4
	**  second  3
	*   first   3
	*** third   5
	**  second  4
	*** third   6
	*** third   7
	**  second  5
	*** third   8
	*** third   9
	*** third   10
	*** finish ***
*/