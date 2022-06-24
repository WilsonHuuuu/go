package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func test1() {
	// r := strings.NewReader("hello world")
	f, _ := os.Open("a.txt")
	defer f.Close()
	r2 := bufio.NewReader(f) //对目标进行封装
	s, _ := r2.ReadString('\n')
	fmt.Printf("s: %v\n", s)
}

func test2() {
	s := strings.NewReader("abcdefghijklmn122346587889521")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	for {
		n, err := br.Read(p) //每次读取10个元素
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p): %v\n", string(p[0:n]))
		}
	}
}

func test3() {
	s := strings.NewReader("abcdefg")
	br := bufio.NewReader(s)

	b, _ := br.ReadByte() //读入了一个元素
	fmt.Printf("b: %c\n", b)

	b, _ = br.ReadByte()
	fmt.Printf("b: %c\n", b)

	br.UnreadByte() //吐出了一个元素，相当于多读入了上一个元素
	b, _ = br.ReadByte()
	fmt.Printf("b: %c\n", b)
}

func test4() {
	s := strings.NewReader("abc\ndef\r\nghijk\r\nlnm")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}

func test5() {
	s := strings.NewReader("abc,def,ghi,jkl")
	br := bufio.NewReader(s)

	w, _ := br.ReadSlice(',')
	fmt.Printf("w: %c\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("w: %c\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("w: %c\n", w)
}

func test6() {
	s := strings.NewReader("abc def ghi jkl")
	br := bufio.NewReader(s)

	w, _ := br.ReadBytes(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadBytes(' ')
	fmt.Printf("%q\n", w)
}

func test7() {
	s := strings.NewReader("abc def ghi jkl")
	br := bufio.NewReader(s)

	w, _ := br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)
}

func test8() {
	s := strings.NewReader("abcdefghijkl")
	br := bufio.NewReader(s)
	b := bytes.NewBuffer(make([]byte, 0))

	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
	defer f.Close()

	br.WriteTo(f) //写入
	fmt.Printf("b: %v\n", b)
}

func test9() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0777)
	defer f.Close()
	w := bufio.NewWriter(f) //具有缓冲会更加高效
	w.WriteString("hello  world")
	w.Flush() //一定要记得刷新
}

func test10() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123456789")
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")
	bw.Flush()
	fmt.Println(b)
	fmt.Println(c)
}

func test11() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)    //写入缓存
	fmt.Println(bw.Available()) //4096
	fmt.Println(bw.Buffered())  //0

	bw.WriteString("abcdefghijklnm")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}

func test12() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123") //读的功能
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	s2, _ := rw.ReadString('\n')
	fmt.Println(string(s2))
	rw.WriteString("asdf") //写的功能
	rw.Flush()
	fmt.Println(b)
}

func test13() {
	s := strings.NewReader("abc def ghi jkl")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords) //以空格作为分隔符进行分割
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

func test14() {
	s := strings.NewReader("hello 世界")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanBytes)
	for bs.Scan() {
		fmt.Printf("%s", bs.Text())
	}
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
	// test11()
	// test12()
	// test13()
	test14()
}
