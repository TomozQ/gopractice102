package main

import (
	"fmt"
	"reflect"
)

// General is all type data.
type General interface {}

// GData is holding General value.
type GData interface {
	Set(nm string, g General) GData
	Print()
}

// NData is structure.
type NData struct {
	Name string
	Data []int
}

// Set is NData method.
func (nd *NData) Set (nm string, g General) GData {
	nd.Name = nm
	// 引数gのKind型を調べる																	0が何型か
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf(0)) {
		nd.Data = g.([]int)
	}
	return nd
}

// Print is NData method
func (nd *NData) Print () {
	fmt.Printf("<<%s>>", nd.Name)
	fmt.Println(nd.Data)
}

// SData is structure
type SData struct {
	Name string
	Data []string
}

// Set is SData method.
func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	// 																										""が何型か
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf("")) {
		sd.Data = g.([]string)
	}
	return sd
}

// Print is SData method.
func (sd *SData) Print () {
	fmt.Printf("* %s [%s] *\n", sd.Name, sd.Data)
}

func main () {
	var data = []GData{}
	data = append(data, new(NData).Set("Taro", 123))
	data = append(data, new(SData).Set("Jiro", []string{"hello!", "bye"}))
	data = append(data, new(NData).Set("Hanako", 98700))
	data = append(data, new(SData).Set("Sachiko", "happy?"))
	for _, ob := range data {
		ob.Print()
	}
}

/*
	型アサーション
	nd.Data = g.(int)
	sd.Data = g.(string)
	gはGeneral型だがこれによりint型やstring型としてDataに設定するようになっている。

	空のインターフェイス型をうまく活用することで保持する値の異なる構造体を同一のインターフェイスとして扱えるようになる。
*/