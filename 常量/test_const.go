package main

import "fmt"

func main() {
	const PI float32 = 3.14

	const PI2 = 3.1415
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("PI2: %v\n", PI2)

	const (
		a = 100
		b = 200
	)
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	const w, h = 200, 300
	fmt.Printf("w: %v\n", w)
	fmt.Printf("h: %v\n", h)
}
