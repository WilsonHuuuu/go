package main

import "fmt"

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break
		}
	}
}

func f2() {
	i := 1
	switch i {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
		fallthrough //如果没有 break ，则当 i：=2 时会输出2，3
	case 3:
		fmt.Println("3")
	}
}

// MYLABEL 是一个标签，可以将标签进行 break
func f3() {
MYLABEL:
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i >= 5 {
			break MYLABEL
		}
	}
}

func main() {
	// f1()
	// f2()
	f3()
}
