package main

import "fmt"

func sumUntil(n int) ([]string, int) {
    sum := 0
    var results []string

    for i := 0; sum < n; i++ {
        // cek jenis angka (i)
        if i == 0 {
            results = append(results, "non")
        } else if i%2 == 0 {
            results = append(results, "genap")
        } else {
            results = append(results, "ganjil")
        }

        // update sum
        sum += i
    }

    return results, sum
}

func main() {
    hasil, total := sumUntil(10)
    fmt.Println("Array hasil:", hasil)
    fmt.Println("Total sum:", total)
}
