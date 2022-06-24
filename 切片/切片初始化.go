package main

import "fmt"

func f1() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:3] //从第0个元素开始，不包含第 3 个元素，只包含 0，1，2 个元素
	fmt.Printf("s2: %v\n", s2)
	s3 := s1[3:] //从第三个元素开始取到最后一个元素
	fmt.Printf("s3: %v\n", s3)
	s4 := s1[2:5]
	fmt.Printf("s4: %v\n", s4)
	s5 := s1[:] //取全部元素
	fmt.Printf("s5: %v\n", s5)

}

//数组同样可以使用切片的规则
func f2() {
	var a1 = [...]int{1, 2, 3, 4, 5, 6}
	a2 := a1[:3]
	fmt.Printf("a2: %v\n", a2)
}

func main() {
	// f1()
	f2()
}
