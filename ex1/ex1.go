package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	p, _ := strconv.ParseUint(os.Args[1], 10, 16)

	rand.Seed(time.Now().UTC().UnixNano())
	n := uint64(rand.Intn(100))

	if n < p {
		fmt.Println("A")
	} else {
		fmt.Println("B")
	}
}
