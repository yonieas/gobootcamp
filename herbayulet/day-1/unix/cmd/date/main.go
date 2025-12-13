package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // mendapatkan waktu sekarang mirip di js new Date()
	fmt.Println(now) 
    fmt.Println(now.Format(time.RFC1123)) // ini seperti mem-format tanggal seperti di format (date-fns) new Date(format(now, "DD, dd MM YYYY HH:mm:ss"))
}