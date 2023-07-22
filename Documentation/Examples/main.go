package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	timeString := t.Format(time.RFC3339)

	fmt.Println("cetak waktu dalam format RFC 3339: ",timeString)
}