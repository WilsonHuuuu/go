package main

import (
	"bytes"
	"fmt"
	"io"
)

// 强制类型转换
func testTrans() {

	var i int = 100
	var b byte = 10
	b = byte(i)
	i = int(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
}

// 是否包含
func testContains() {
	s := "duoke360.com"
	b := []byte(s)
	b1 := []byte("duoke360")
	b2 := []byte("duoke720")

	b3 := bytes.Contains(b, b1)
	fmt.Printf("b3: %v\n", b3)
	b3 = bytes.Contains(b, b2)
	fmt.Printf("b3: %v\n", b3)
}

// 计数
func testCount() {
	s := []byte("helloooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1))
	fmt.Println(bytes.Count(s, sep2))
	fmt.Println(bytes.Count(s, sep3))

}

// 重复内容
func testRepeat() {
	b := []byte("hello")
	fmt.Println(string(bytes.Repeat(b, 1)))
	fmt.Println(string(bytes.Repeat(b, 3)))
}

// 替换内容
func testReplace() {
	s := []byte("hello world")
	old := []byte("o")
	new := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, new, 0)))  //替换零个
	fmt.Println(string(bytes.Replace(s, old, new, 1)))  //替换一个
	fmt.Println(string(bytes.Replace(s, old, new, 2)))  //替换两个
	fmt.Println(string(bytes.Replace(s, old, new, -1))) //全部替换
}

// 可以转换汉字
func testRunes() {
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println("转换前字符串的长度：", len(s))
	fmt.Println("转换后字符串的长度：", len(r))
}

func testJoin() {
	s2 := [][]byte{[]byte("你好"), []byte("世界")} //[][]byte是二维切片，里面初始化两个切片
	sep4 := []byte(",")                        //连接切片的是个 ，
	fmt.Println(string(bytes.Join(s2, sep4)))  //输出：你好，世界
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5))) //输出：你好#世界
}

func testReader() {
	data := "123456789"
	//通过 []byte 创建 Reader
	re := bytes.NewReader([]byte(data))
	//返回读取部分的长度
	fmt.Println("re.len :", re.Len())
	// 返回底层数据总长度
	fmt.Println("re size :", re.Size())
	fmt.Println("------------")

	buf := make([]byte, 2)
	for {
		//读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	fmt.Println("------------")

	//设置偏移量
	re.Seek(0, 0)
	for {
		//一个字节一个字节的读
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}
	fmt.Println("-----------------")

	re.Seek(0, 0)
	off := int64(0)
	for {
		//指定偏移量读取
		n, err := re.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}

func testBuffer() {
	var b bytes.Buffer
	fmt.Printf("b: %v\n", b)
	var b1 = bytes.NewBufferString("hello")
	fmt.Printf("b1: %v\n", b1)
}

func testBuffer2() {
	var b bytes.Buffer
	n, _ := b.WriteString("hello")
	fmt.Printf("n: %v\n", n)
	fmt.Printf("b: %v\n", b.Bytes())
	fmt.Printf("b: %v\n", string(b.Bytes()))

}

// 读取
func testBuffer3() {
	var b = bytes.NewBufferString("hello world")
	b1 := make([]byte, 2)
	for {
		n, err := b.Read(b1)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(b1): %v\n", string(b1[0:n]))
	}

}

func main() {
	// testReader()
	testBuffer3()
}
