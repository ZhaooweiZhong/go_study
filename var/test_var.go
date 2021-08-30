package main

/*
	四种变量名称
	声明全局变量，前三种不变，
	方法四只能在函数体内使用
*/
import (
	"fmt"
)

func main() {
	//方法一：声明一个变量，默认值是0
	var a int
	fmt.Println(a)

	//方法二：声明一个变量，初始化一个值
	var b int = 1000
	fmt.Println(b)

	//方法三：初始化的时候省去类型，自动匹配当前变量的数据类型
	var c = 1000
	fmt.Println("c=", c)
	fmt.Printf("Type of c = %T \n", c)

	//方法四：省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("e = %T", e)
}
