package bluecore

func Lesson(cfg Config) error {

	return nil
}

// Structs
type myStruct struct {
}

type MyStruct struct {
	Impl  myStruct
	privA myStruct
	privB myStruct
	A     string
	B     interface{}
}

// New() type methods help to instanciate any private variables
// held within
func NewMyStruct() MyStruct {
	return MyStruct{
		Impl: myStruct{},
		A:    "asdf",
	}
}

/*
	class MyStruct:
		class myStruct:
			def Do():
				pass
		def DoStuff(a):
			pass
*/
// func (obj) name(params) return type {}
func (m MyStruct) DoStuff(data string) {
	// Can operate on private variables declared within
	// MyStruct
	m.A = "something else"

	return
}

// This function operates on MyStruct, but isn't a method OF MyStruct
func DoStuff(m MyStruct, data string) {
}

type BStruct struct {
}

func (b BStruct) DoStuff(s string) {
}

// Interfaces
type IStruct interface {
	DoStuff(a string)
}

func OperateOnIface(inst IStruct) {
	inst.DoStuff("asdfads")
}

// Struct Methods

// Pointers
func (b *BStruct) WithPointers(s string) {

}
