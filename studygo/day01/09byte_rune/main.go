package main

import "fmt"

func main() {

	//字符串修改，字符串是不能修改的变量，只能先转换为[]byte活[]tune，再修改，
	//再转换为string
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := '红' //rune(int32)                 
	c2 := "红" //string
	fmt.Printf("%T,%T",c1,c2)

}
