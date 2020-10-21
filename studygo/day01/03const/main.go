package main

import "fmt"
//常量
//定义了常量之后不能修改
const pi = 3.1415926

const(
	statusOk = 200
	notFound = 404
)
const(
	n1 = 100
	n2
	n3
	//批量声明常量时，如果某一行声明后没有赋值，默认和上一行一样
)

//iota
//iota是go语言的常量计数器，只能在常量的表达式中使用
//iota在const关键字出现时被重置为0，const中每新增一行常量声明，iota自增1
const(
	a1 = iota 
	a2     //a2 = iota
	a3 
)


//跳过某些值
const(
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)

//插队
const(
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

//多个常量声明在一行
const(
	d1,d2 = iota+1,iota+2 //1,2
	d3,d4 = iota+1,iota+2 //2,3
)

//定义数量级
const(
	_ = iota
	KB = 1<<(10 * iota) //1024
	

)



func main(){
	//pi =123
	fmt.Println("n1:",n1)
	fmt.Println("n2:",n2)
	fmt.Println("n3:",n3)

	fmt.Println("n1:",b1)
	fmt.Println("n2:",b2)
	fmt.Println("n3:",b3)

	fmt.Println("c1:",c1)
	fmt.Println("c2:",c2)
	fmt.Println("c3:",c3)
	fmt.Println("c4:",c4)
}