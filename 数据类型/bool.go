package main

import "fmt"

func main() {

	// 在if中不能用0或非0表示条件

	age := 18
	gender := "男"

	if age >= 18 && gender == "男" {
		fmt.Println("成年男性")
	}

	// count := 10
	// for i := 0; i < count; i++ {
	// 	// fmt.Printf("i=%v\n", i)
	// 	fmt.Printf("i: %v\n", i)
	// }

	// if 语句
	// age := 16
	// r := age >= 18
	// fmt.Printf("r: %v\n", r)
	// if r {
	// 	fmt.Println("成年")
	// } else {
	// 	fmt.Println("未成年")
	// }

	// var b1 bool = true
	// var b2 bool = false

	// var b3 = true
	// var b4 = false

	// b5 := true
	// b6 := false

	// fmt.Printf("b1: %v\n", b1)
	// fmt.Printf("b2: %v\n", b2)
	// fmt.Printf("b3: %v\n", b3)
	// fmt.Printf("b4: %v\n", b4)
	// fmt.Printf("b5: %v\n", b5)
	// fmt.Printf("b6: %v\n", b6)
}
