// An implementation of Conway's Game of Life.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	w, h := GetDims()
	l := NewLife(w, h-1)

	for {
		l.Step()
		fmt.Print("\033[H\033[2J", l) // Clear screen and print field.
		time.Sleep(time.Second / 30)
	}
}
