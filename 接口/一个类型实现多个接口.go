package main

import "fmt"

type Player interface {
	playMusic()
}

type Video interface {
	playVideo()
}

type Mobile struct {
}

func (m Mobile) playMusic() {
	fmt.Println("play music")
}

func (m Mobile) playVideo() {
	fmt.Println("play video")
}

func main() {
	m := Mobile{}
	m.playMusic()
	m.playVideo()
}
