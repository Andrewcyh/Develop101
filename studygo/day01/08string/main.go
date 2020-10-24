package main
//字符串
import "fmt"
func main(){
	//\为转义符
	s3 :=
	//多行字符串
	s2 := `
	round1
	ko!
	win!
	`
	fmt.Println(s2)

	//字符串相关操作
	fmt.Println(len(s2))

	//字符串拼接
	name :="理想"
	world := "dsb"
	
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s+%s",name,world)

	fmt.Printf("%s%s",name,world)
	fmt.Println(ss1)

	//分割
	ret := strings.Split(s3,"\\")
	fmt.Println(ret)
	//包含
}