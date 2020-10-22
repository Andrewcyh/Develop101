package main

import "fmt"

//常量
//定义后不能改变
const pi = 3.1415926

//批量声明常量
const (
	statusOk = 200
	notFound = 404
)

//批量声明常量，某一行开始没有等号复制，默认于上一行相等
const (
	n1 = 100
	n2
	n3
)

//iota  常量计数器
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)

const (
	c1 = iota //0
	c2 = 100
	c3 //2
	c4 //3
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1 = 1,d2 = 2
	d3, d4 = iota + 1, iota + 2 //d3 = 2,d4 = 3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

//八进制&十六进制

func main() {
	var i1 = 101
	fmt.Printf("%d", i1)
	fmt.Printf("%b", i1) //二进制
	fmt.Printf("%o", i1) //八进制
	fmt.Printf("%x", i1) //十六进制

	i2 := 077    //八进制
	i3 := 0x6677 //十六进制

	//查看变量类型
	fmt.Printf("%T", i3)
	//int8类型变量声明
	i4 := int8(9)

}
