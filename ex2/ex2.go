package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Field wraps a 2 dimension boolean array
type Field struct {
	lines [][]bool
}

func generateField(nbLines uint, nbRows uint, density uint) *Field {
	rand.Seed(time.Now().UTC().UnixNano())
	lines := make([][]bool, nbLines)

	for i := range lines {
		lines[i] = make([]bool, nbRows)

		for j := range lines[i] {
			n := uint(rand.Intn(100))

			if n < density {
				lines[i][j] = true
			} else {
				lines[i][j] = false
			}
		}
	}

	return &Field{
		lines,
	}
}

func (field *Field) String() string {
	var buf bytes.Buffer

	for _, line := range field.lines {
		for _, cell := range line {
			if cell {
				buf.WriteByte('0')
			} else {
				buf.WriteByte(' ')
			}
		}

		buf.WriteByte('\n')
	}

	return buf.String()
}

func main() {
	nbLines, _ := strconv.ParseUint(os.Args[1], 10, 32)
	nbRows, _ := strconv.ParseUint(os.Args[2], 10, 32)
	density, _ := strconv.ParseUint(os.Args[3], 10, 32)

	field := generateField(uint(nbLines), uint(nbRows), uint(density))

	fmt.Print(field)
}
