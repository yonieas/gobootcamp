package main

import (
	"fmt"
	"os"
)

func main() {
	// harus ada 1 argumen: cat <file>
	if len(os.Args) != 2 { //args[0] itu yang go run & args[1] itu yang nama.txt atau cat (binary)
		fmt.Println("usage: cat <file>") // ini jika dijalankan hanya go run main.go
		return
	}

	filename := os.Args[1] // nah ini untuk yang name.txt itu

	// baca seluruh isi file
	data, err := os.ReadFile(filename) // harus di build dulu file nya dengan go build -o cat main.go (harus di directory yang dimana file tersebut ada)
	if err != nil {
		fmt.Println("error weyy:", err) // jika argumen kedua nya tidak cocok
		return
	}

	// print isi file
	fmt.Printf("%s\n", data) // jika ada data nya atau file yang dibuka

}
