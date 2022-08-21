package main

import (
	"time"
	"fmt"
	"sync"
	"strconv"
)

// SrData is structure.
type SrData struct {
	msg string
	mux	sync.Mutex
}

func main () {
	sd := SrData{msg: "Start"}	// Mutexは初期状態のままで、新たに値などを設定する必要はない。
	prmsg := func(nm string, n int) {
		fmt.Println(nm, sd.msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	main := func(n int) {
		const nm string = "*main"
		sd.mux.Lock()
		for i := 0; i < 5; i++ {
			sd.msg += " m" + strconv.Itoa(i)
			prmsg(nm, 100)
		}
		sd.mux.Unlock()
	}

	hello := func (n int) {
		const nm string = "hello"
		// sd.mux.Lock()
		for i := 0; i < 5; i++ {
			sd.msg += " h" + strconv.Itoa(i)
			prmsg(nm, n)
		}
		sd.mux.Unlock()
	}

	go main(100)
	go hello(50)
	time.Sleep(3 * time.Second)
}

/*
	出力結果
	hello Start h0
	hello Start h0 h1
	hello Start h0 h1 h2
	hello Start h0 h1 h2 h3
	hello Start h0 h1 h2 h3 h4
	*main Start h0 h1 h2 h3 h4 m0
	*main Start h0 h1 h2 h3 h4 m0 m1
	*main Start h0 h1 h2 h3 h4 m0 m1 m2
	*main Start h0 h1 h2 h3 h4 m0 m1 m2 m3
	*main Start h0 h1 h2 h3 h4 m0 m1 m2 m3 m4
*/

/*
	main関数,hello関数ともスレッドをロックしてから繰り返し処理を実行
	繰り返しを抜けたらアンロック
	他スレッドがsdにアクセスできなくなるため、もう一方の処理がアンロックされるまで待ち続けることになる。
	
	スレッドをロックすると、アンロックされるまでの間たスレッドはペンディング状態となる
	別スレッドで実行している処理を全てロックしてしまうと、結局並行処理ではなく「複数のスレッドが順に実行される」という逐次処理になってしまう。
	ロックは「この処理だけは外部からアクセスされては困る」という必要最低限の範囲に絞って行うようにする。
*/