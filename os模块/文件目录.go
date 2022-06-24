package main

import (
	"fmt"
	"os"
)

// 创建文件
func creatFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func makeDir() {
	err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 创建目录中的目录
	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 删除目录或者文件
func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.RemoveAll("a")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

func wd() {

	// 获取目录
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	// 修改目录
	os.Chdir("d:/")
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)

	// 获得当前的临时目录
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// 重命名目录
func rename() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func readFile() {
	b, _ := os.ReadFile("test2.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

// 写文件
func writeFile() {
	os.WriteFile("test2.txt", []byte("hello"), os.ModePerm)
}

func main() {
	// creatFile()
	// makeDir()
	// remove()
	// wd()
	// rename()
	// readFile()
	writeFile()
}
