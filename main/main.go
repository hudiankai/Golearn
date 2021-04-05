package main

import (
	"Golearn/Learn"
	"fmt"
)


func main()  {
	fmt.Println("ada")
	Liming := Learn.Student{ 	Age: 12, 	Sex: "Man", 	Num: 1, }
	Liming.RML.RLock()
	fmt.Println(Liming)
	Learn.ReadandRead(&Liming)
	defer Liming.RML.RUnlock()
	//Learn.LockandRead(&Liming)
}

