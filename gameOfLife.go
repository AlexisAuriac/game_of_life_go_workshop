// An implementation of Conway's Game of Life.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	args, err := ParseArgs()

	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	w, h := GetDims()
	l := NewLife(w, h-1, args.colorAlive, args.colorDead)

	for {
		l.Step()
		fmt.Print("\033[H\033[2J", l) // Clear screen and print field.
		time.Sleep(time.Second / time.Duration(args.speed))
	}
}
