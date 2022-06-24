package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"
	n := 3
	m := 5
	fmt.Println(str[n])   //获取索引位置为 n 的原始字节
	fmt.Println(str[n:m]) //截取字符串 n 到 m-1 的字符串
	fmt.Println(str[n:])  //截取字符串 n 到 len(s)-1 的字符串
	fmt.Println(str[:m])  //截取字符串 0 到 m-1 的字符串

	s := "hello world"
	a := 2
	b := 7
	fmt.Printf("s[a]: %c\n", s[a])
	fmt.Printf("s[a:b]: %v\n", s[a:b]) //取值取到 b 的前一位
	fmt.Printf("s[a:]: %v\n", s[a:])
	fmt.Printf("s[:b]: %v\n", s[:b])

	// 字符串长度
	S := "hello world"
	fmt.Printf("len(S): %v\n", len(S))
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))               //分割字符串
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello")) //是否包含
	fmt.Printf("strings.ToLower(s): %v\n", strings.ToLower(s))                       //转换小写
	fmt.Printf("strings.ToUpper(s): %v\n", strings.ToUpper(s))                       //转换大写

	fmt.Printf("strings.HasPrefix(s, \"Hello\"): %v\n", strings.HasPrefix(s, "Hello")) //是否以 “” 里面的内容开头（区分大小写）
	fmt.Printf("strings.HasSuffix(s, \"world\"): %v\n", strings.HasSuffix(s, "world")) //是否以 “” 里面的内容结尾

	fmt.Printf("strings.Index(s, \"ll\"): %v\n", strings.Index(s, "ll")) //索引 “” 里面内容的位置
}
