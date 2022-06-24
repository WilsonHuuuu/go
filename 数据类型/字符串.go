package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var s string = "hello"
	var s1 = "hello"
	s3 := "hello"

	fmt.Printf("s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s3: %v\n", s3)

	s4 := `
	line 1
	line 2
	line 3
	` //多行字符串
	fmt.Printf("s4: %v\n", s4)

	// 字符串连接
	S1 := "tom"
	S2 := "20"
	msg := S1 + S2
	fmt.Printf("msg: %v\n", msg)

	name := "tom"
	age := "20"

	S := strings.Join([]string{name, age}, ",")
	fmt.Printf("s: %v\n", S)

	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(",")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())
}
