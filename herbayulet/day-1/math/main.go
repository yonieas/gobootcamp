package main

import (
	"fmt"
	"math"
	"math/rand"
)

func randomInt(n int) int {
	return rand.Intn(n) // ini melakukan print nilai random berdasarkan package rand 
}

func akar(n float64) float64 { //walopun paramater angka bulat, c: 25 tp method sqrt ini selalu menerapkan jika tidak angka bulat makanye diberi float64 karena digitnya lebih panjang, dan ketika di kasih int biasa dia ada warning
	return math.Sqrt(n)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9 
	y = sum - x
	return
}

func main() {
	fmt.Println("My favorite number is", randomInt(10)) // dan dipanggil disini func randomItn
	fmt.Printf("Now you have %g problems.\n", akar(25)) // dan dipanggil disini func akar dengan %g untuk menyederhanakan output cmiiw
	fmt.Println(split(10)) // dan dipanggil disini func split
}
