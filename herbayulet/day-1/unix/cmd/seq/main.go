package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// harus ada 1 argumen: seq N
	if len(os.Args) != 2 {
		fmt.Println("usage: seq N") // ini bakal di print jika tidak ada argumen yang kita input
		return
	}

	// konversi argumen ke integer
	n, err := strconv.Atoi(os.Args[1]) // ini mengubah stdin kita yang mana string menjadi integer
	// fmt.Println(n, "ini n") // untuk menjalakan ini ketika run masuka argumen nya contoh: go run main.go 10 maka n nya itu 10 dan bakal return looping dibawah
	// fmt.Println(err, "ini err") // untuk menjalakan ini ketika run masukan argumen nya contoh: go run main.go "herbayu" maka n nya itu 10 dan bakal return kondisi dibawah
	if err != nil {
		fmt.Println("invalid number:", os.Args[1])
		return
	}

	// print angka dari 1 sampai N
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
