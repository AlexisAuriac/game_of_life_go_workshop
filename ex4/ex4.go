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

func generateField(height uint, width uint, density uint) *Field {
	rand.Seed(time.Now().UTC().UnixNano())
	lines := make([][]bool, height)

	for y := range lines {
		lines[y] = make([]bool, width)

		for x := range lines[y] {
			n := uint(rand.Intn(100))

			if n < density {
				lines[y][x] = true
			} else {
				lines[y][x] = false
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

func (field *Field) copy(other *Field) {
	field.lines = make([][]bool, len(other.lines))

	for y := range field.lines {
		field.lines[y] = make([]bool, len(other.lines[y]))

		for x := range field.lines[y] {
			field.lines[y][x] = other.lines[y][x]
		}
	}
}

func cellIsAlive(field *Field, x int, y int) bool {
	if y < 0 || x < 0 || y >= len(field.lines) || x >= len(field.lines[y]) {
		return false
	}

	return field.lines[y][x]
}

func countLiveNeighbours(field *Field, x int, y int) uint {
	live := uint(0)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			} else if cellIsAlive(field, x+i, y+j) {
				live++
			}
		}
	}

	return live
}

func conway(field *Field) *Field {
	newField := new(Field)
	newField.copy(field)

	for y := range field.lines {
		for x := range field.lines[y] {
			live := countLiveNeighbours(field, x, y)

			if field.lines[y][x] && live == 2 {
				continue
			} else if live == 3 {
				newField.lines[y][x] = true
			} else {
				newField.lines[y][x] = false
			}
		}
	}

	return newField
}

func main() {
	nbLines, _ := strconv.ParseUint(os.Args[1], 10, 32)
	nbRows, _ := strconv.ParseUint(os.Args[2], 10, 32)
	density, _ := strconv.ParseUint(os.Args[3], 10, 32)

	field := generateField(uint(nbLines), uint(nbRows), uint(density))
	fmt.Print(field)

	for {
		time.Sleep(time.Second / time.Duration(10))
		for i := uint64(0); i < nbRows; i++ {
			fmt.Print("-")
		}
		fmt.Println()

		field = conway(field)
		fmt.Print(field)
	}
}
