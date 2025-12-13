package main

import "fmt"

var c, python, java bool //declare type data yang mana default boolean adalah false

func initializers() (bool, bool, string){
	var d, pythons, javas = true, false, "no!"
	return d, pythons, javas
}

func main() {
	var i int // jika tidak ada value default int adalah 0
	fmt.Println(i, c, python, java)
	fmt.Println(initializers())
}
