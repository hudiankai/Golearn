package Learn

import (
	"fmt"
	"time"
)

func test(j int)  {
	for i := 0;i<5;i++{
		time.Sleep(1*time.Millisecond)
		fmt.Println(i+j)
	}

}


func CH()  {
	go test(10)
	test(4)
}
