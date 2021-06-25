//Init函数是什么时候执行
package Mianshi

import "fmt"

const A = 10

var B= 100

func init()  {
	fmt.Println("Init2",A,B)
}

func init()  {
	fmt.Println("Init1",A,B)
}

//func main()  {
//	fmt.Println("Main",a,b)
//}

/*
输出如下：
Init2 10 100
Init1 10 100
Main 10 100
从输出可以看出，init函数的执行在Main函数之前，并且同一个包甚至是同一个源文件可包含多个init函数，init函数的执行顺序没有保证
init函数是Go程序初始化的一部分，没有输入参数和返回值，不能被其他函数调用，
执行顺序：import > const > var > init() > main()
（1）由runtime按照依赖关系（没有依赖的包最先初始化，Learn包中调用了“B“，固其init函数后执行）初始化每个导入（import）的包，
（2）每个包先初始化包作用域内的常量（const）和变量
（3）执行包中的init函数，同一个包可多个，但执行顺序没有保证
（4）执行main函数
*/
