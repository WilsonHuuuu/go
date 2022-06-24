package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//原子变量就是保护协程的操作
var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

/*var i = 100

var lock sync.Mutex

func add() {
	lock.Lock()
	i++
	lock.Unlock()
}

func sub() {
	lock.Lock()
	i--
	lock.Unlock()
}*/

func main() {

	for i := 0; i < 100; i++ {
		go add()
		go sub()
	}

	time.Sleep(time.Second)

	fmt.Printf("i: %v\n", i)
}
