package main

import (
	"fmt"
)

type Printer interface {
	Print(test int, test2 int)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) Print(test int, test2 int) {
	fmt.Println("helo")
}

func (r Rectangle) Test() {
	fmt.Println("test")
}

func testHe() {
	var a interface{} = 1

	i := a
	j := a.(int)

	println(i)
	println(j)

	var b int
	b = 1

	b_pointer := &b
	b_none_pointer := *b_pointer

	println(b_pointer, b_none_pointer)

	var tester Printer
	r := Rectangle{1, 2}
	tester = r
	tester.Print(1, 2)

	var r1 []byte
	str := "ㄱ"
	r1 = []byte(str)
	println(string(r1[:]))
	println(r1)
}

func test2() {
	slice := []int{0, 10, 20, 30}
	slice2 := []int{}

	copyslice := make([]int, len(slice))
	test := copy(copyslice, slice)

	slice2 = slice[test:]

	println(test)
	println(slice[3])
	println(slice)
	println(slice2)
}

func main() {
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }
}
