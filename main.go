package main

import (
	"fmt"
)

/*
	interface
	その構造体に必要なメソッドが必ず用意されていることを保証する
	構造体に付加するメソッドをまとめたもの
	typeとして作成する
	type 名前 interface {
		メソッドA
		メソッドB
		'''必要なだけ用意'''
	}
*/

// Data is interface
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

// Mydata is struct
type Mydata struct {
	Name string
	Data []int
}

// Initial is init method
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// PrintData is println all data.
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// Check is method
func (md *Mydata) Check() {
	fmt.Printf("check! [%s]", md.Name)
}

func main () {
	var ob Mydata = Mydata{}
	// var ob Data = new(Mydata)　	Data型のため’Check’メソッドは使用できない
	/*
		new関数は、これ自体は特定の方の値を作成するものではない。
		代入する変数の方に合わせて値は扱われる。
		ここでは
		var ob Data
		と変数が定義されているため、そこにnewで代入した値はData型の値と判断され、そう扱われるようになる。

		逆に
		var ob Data = Mydata{}
		これは「Mydataの値をDate型の変数に代入しようとしている」としてエラーになる。

		さらに
		var ob Data = Data()
		や
		var ob Data = Data{}
		これはいずれもエラーになる。
		Dataはインターフェイスであり、一般的な構造体などとは違うもの。
		インターフェイスは直接値を作成することはできない。
		だからこそnewを使ってMydataを作成し、それをDataと推論させて変数に代入させている。
	*/
	ob.Initial("Sachiko", []int{55,66,77})
	// ob.PrintData()
	ob.Check()
}


/*
	関連メモ
	-----------------------------------
	number := 123
	pointer := &number
	value := *pointer

	number -> int型の変数
	pointer -> ポインタ型の変数（numberのアドレス）
	value -> pointer(numberのアドレス)に入っている値
	-----------------------------------
*/

/*
	n := 123  		...int型の変数n
	change2(&n)		...change2にnのポインタを渡す

	func change2 (n *int) {		...int型の値が入ったポインタ型の引数n
		*n *= 2									...ポインタの値に2をかける
	}
*/