package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	var a = 'a'
	var b = "a"
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	var s = "Hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

}
