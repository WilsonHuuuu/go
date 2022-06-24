package main

import (
	"fmt"
	"os"
)

func openClose() {
	// 方法1
	f, err := os.Open("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
		f.Close()
	}

	// 方法2
	f2, err2 := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755) //等价于 os.Creat
	if err2 != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
		f.Close()
	}
}

// 创建文件
func creatFile() {
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	// （）里面第一个参数：目录默认；第二个参数：文件名前缀
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// 读文件
func readOps() {
	// 循环读取文件
	/*f, _ := os.Open("a.txt")
	for {

		buf := make([]byte, 3) //缓冲区
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}

	f.Close()*/

	// 从某一个开始读取
	f, _ := os.Open("a.txt")
	buf := make([]byte, 3)
	n, _ := f.ReadAt(buf, 3)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))

	// 读取目录
	de, _ := os.ReadDir("a")
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}

	// 定位
	os.Open("a.txt")
	f.Seek(3, 0) //Seek 寻找位置
	buf1 := make([]byte, 10)
	n2, _ := f.Read(buf1)
	fmt.Printf("n2: %v\n", n2)
	fmt.Printf("string(buf1): %v\n", string(buf1))
	f.Close()
}

func main() {
	// openClose()
	readOps()
}
