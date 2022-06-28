package another

import (
	"fmt"
	"github.com/tj10200/golang-class/class_1/01_structs_interfaces_and_mocks/pkg/bluecore"
)

func A() {
	// Can operate on public variables
	// of MyStruct only
	m := bluecore.MyStruct{}
	m.DoStuff("asdfads")
}

func B() {
	m := bluecore.MyStruct{}
	bluecore.OperateOnIface(m)
	b := bluecore.BStruct{}
	bluecore.OperateOnIface(b)
}

func C() {
	data := make([]byte, 20)
	data2 := data[1:5]
	data3 := data[10:]

	for i, _ := range data2 {
		data2[i] = byte('a' + i)
	}

	for i, _ := range data3 {
		data3[i] = byte('a' + i)
	}

	fmt.Printf("D1 (%s)\n D2 (%s) D3 (%d)",
		string(data),
		string(data2),
		string(data3))
}
