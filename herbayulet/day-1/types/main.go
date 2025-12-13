package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var ( // membuat variabel dengan cara lebih rapih
	ToBe   bool       = false
	MaxInt uint64     = 1<<2 - 1
	z      complex128 = cmplx.Sqrt(-5 + 10)
)

func conversion() (int, int, uint) {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) 
	var z uint = uint(f) //mengkorvensi hasil dari f yang float menjadi integer positif
	return x, y, z
}

func main() {
	fmt.Printf("Type: %T /* <= deklrasi type data kalo di js typeof */ Value: %v /*  <= deklrasi value nya */\n", ToBe, ToBe) // argumen 1 untuk type data dan kedua untuk print value nya
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println(conversion())
}
