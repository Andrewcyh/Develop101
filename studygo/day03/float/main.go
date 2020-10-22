package main

import (
	"fmt"
	"math"
)

func main() {
	math.MaxFloat32
	fmt.Printf("%d", math.MaxFloat32)
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认都是float64类型
	f2 := float32(1.23456)
	//f1 = f2 // float32不能直接赋值给float64

	//布尔值
	//默认为false
	//不能强制转换为整数
	//
	var b1 bool
	b2 := true
	fmt.Printf("%T,%v", b2, b2) //%v  所有类型都可以打印

	//查看类型
	fmt.Printf("%T", b1)

}
