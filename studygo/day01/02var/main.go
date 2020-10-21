package main

import "fmt"
//推荐使用小驼峰式命名
var studentName string
var StudentName string
var student_name string
//批量声明
var (
	name string 
	age int
	isOk bool
)

func main(){
	name = "理想"
	age = 16
	isOk = true

	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s\n",name)
	fmt.Println(age)

	//简短变量声明：只能用在函数里
	s3 :="hahaha"
	//匿名变量_,不占用内存,不占用命名空间

	/*
	1、_多用于接受不使用的值
	2、函数外的语句必须以关键字开始，如func,var
	3、一个作用域不能重复声明同名的变量

}
