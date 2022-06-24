package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	r := strings.NewReader("hello world")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func reader() {
	r := strings.NewReader("hello world")
	buf := make([]byte, 20)
	r.Read(buf)
	fmt.Printf("string(buf): %v\n", string(buf))
}

func main() {
	// reader()
	testCopy()
}
