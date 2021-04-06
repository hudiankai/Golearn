package Learn

import (
	"fmt"
	"sync"
)
/*main函数中的代码
Liming := Learn.Student{ 	Age: 12, 	Sex: "Man", 	Num: 1, }
Liming.RML.RLock()
fmt.Println(Liming)
Learn.ReadandRead(&Liming)
defer Liming.RML.RUnlock()
 */

type Student struct {
	Age int
	Sex string
	Num int
	RML sync.RWMutex    //读写锁中也可以使用Lock
	L   sync.Mutex      //互斥锁中没有RLock
}

//互斥锁内使用读写锁
func LockandRead(S *Student)  {
	S.RML.Lock()
	defer S.RML.Unlock()
	fmt.Println(S)

	S.RML.RLock()
	fmt.Println(S)
	defer S.RML.RUnlock()
}

//读写锁内使用读写锁(四个函数只有这一个是没有错误的,只有读锁可以互相嵌套)
func ReadandRead(s *Student)  {
	s.RML.RLock()
	defer s.RML.RUnlock()
	fmt.Println(s)

	s.RML.RLock()
	fmt.Println(s)
	defer s.RML.RUnlock()
}

//读写锁内使用互斥锁
func ReadandLock(s *Student)  {
	s.RML.RLock()
	defer s.RML.RUnlock()
	fmt.Println(s)

	s.RML.Lock()
	fmt.Println(s)
	defer s.RML.Unlock()
}

//互斥锁内使用互斥锁
func LockandLock(s *Student)  {
	s.RML.Lock()
	fmt.Println(s)
	defer s.RML.Unlock()

	s.RML.Lock()
	fmt.Println(s)
	defer s.RML.Unlock()
}



