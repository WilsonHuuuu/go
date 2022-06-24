package main

import (
	"fmt"
	"sync/atomic"
)

func add_sub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i: %v\n", i)
	atomic.AddInt32(&i, -1)
	fmt.Printf("i: %v\n", i)

	var j int64 = 200
	atomic.AddInt64(&j, 1)
	fmt.Printf("j: %v\n", j)
}

func load_store() {
	var i int32 = 100
	atomic.LoadInt32(&i) //read
	fmt.Printf("i: %v\n", i)

	atomic.StoreInt32(&i, 200)
	fmt.Printf("i: %v\n", i)
}

// 这个是其他操作的根基
func cas() {
	var i int32 = 100
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("i: %v\n", i)
}

func main() {

}
