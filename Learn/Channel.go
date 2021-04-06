package Learn

import (
	"fmt"
	"time"
)

//遍历通道,可直接调用执行
func Rangechan()  {
	ch := make(chan int,10)//因为是先go协程，后遍历，此处无缓冲通道也行
	//go Closechan(ch)
	//上文go + Closechan函数与下文中的go func(){}()是相同的作用
	go func() {
		for i:=0;i<10;i++ {
			ch <- i
		}
		close(ch)
	}()
	//便利通道，只有value，没有key，通道遵循先进先出
	for num := range ch{
		fmt.Println(num)
	}
}

//关闭通道
func Closechan(ch chan int)  {
	//ch := make(chan int)
	for i:=0;i<10;i++ {
		ch <- i
	}
	//close函数是一个内建函数， 用来关闭channel，这个channel要么是双向的， 要么是只写的（chan<- Type）。
	//这个方法应该只由发送者调用， 而不是接收者。
	//当最后一个发送的值都被接收者从关闭的channel(下简称为c)中接收时,
	//接下来所有接收的值都会非阻塞直接成功，返回channel元素的零值。
	close(ch)
}





//对无缓冲通道进行写入
func Set(ch chan Student,num Student)  {
	//ch.Rl.RLock()
	fmt.Print("写入",num)
	ch <- num
	//defer ch.Rl.RUnlock()

}

//取出无缓冲通道中的数据
func Get(ch chan Student)  {
	//ch.Rl.RLock()
	a := <- ch
	fmt.Print("取出",a)
	//defer ch.Rl.RUnlock()
}

/*用于测试无缓冲通道的main函数
func main()  {
	s := Learn.Student{
		Age: 12,
		Sex: "Man",
		Num: 1,
	}

	var ch chan Learn.Student//声明ch为通道
	ch  = make(chan Learn.Student)//初始化为无缓冲通道，也可以初始化为有传冲通道
	//以上两行代码与下一行注释的代码是相同作用
	//ch := make(chan Learn.Student)
	arr := []Learn.Student{s,s,s,s,s}
	for _,v := range arr{
		go Learn.Set(ch,v)
		go Learn.Get(ch)
	}
	//延时的作用在于给两个go协程运行时间，不然主线程结束运行的话，协程自动结束了运行
	time.Sleep(19*time.Millisecond)

}
 */


func test(j int)  {
	for i := 0;i<5;i++{
		time.Sleep(1*time.Millisecond)
		fmt.Println(i+j)
	}

}

//go对应的协程test函数，会因为主线程（CH函数）已经运行结束，故不再运行
func Xiancheng()  {
	test(20)
	go test(27)//
}

//理论上先运行协程go test函数，故一般情况协程函数与线程函数test都会运行完成
func Xiecheng()  {
	go test(4)
	test(10)
}


/*
func main()  {
	fmt.Println("线程")
	Learn.Xiancheng()
	fmt.Println("协程")
	Learn.Xiecheng()
}
 */