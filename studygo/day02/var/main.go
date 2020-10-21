package main

import "fmt"

//声明变量(变量名命名尽量使用小驼峰式命名   studentName)
//var name string
//var age int
//var isOk bool

//批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "lixiang"
	age = 16
	isOk = true
	//Go语言中变量声明后必须使用，否则无法通过编译
	fmt.Print(age)                 //在终端中打印输出内容
	fmt.Println(isOk)              //打印输出内容后自动添加一个换行符
	fmt.Printf("name is:%s", name) //%s是占位符

	//var aa = 66 //声明变量同时赋值
	//var aaa string = "qwertyuiop"

	//s3 := 77 //简短声明
	//匿名变量用"_"表示
	//同一个作用域下不能重复命名同名变量
}
