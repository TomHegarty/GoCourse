package main

import "fmt"

type str string

func (text str) log() {
	fmt.Println(text)
}

func customTypes() {
	var name str = "Tom"

	name.log()
}
