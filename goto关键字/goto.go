package main

import "fmt"

// 跳出单层循环(跳转到标签里的内容)
func f1() {
	i := 1
	if i >= 2 {
		fmt.Println("2")
	} else {
		goto END
	}
END:
	fmt.Println("END....")
}

// 跳出双层循环
func f2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i >= 2 && j >= 2 {
				goto END
			}
			fmt.Printf("%v,%v\n", i, j)
		}
	}
END:
	fmt.Println("END...")
}

func main() {
	// f1()
	f2()
}
