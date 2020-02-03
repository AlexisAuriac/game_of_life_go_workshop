package main

import (
    "bytes"
    "math/rand"
)

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
    a, b                  *Field
    w, h                  int
    colorAlive, colorDead string
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(w, h int, colorAlive, colorDead string) *Life {
    a := NewField(w, h)
    for i := 0; i < (w * h / 4); i++ {
        a.Set(rand.Intn(w), rand.Intn(h), true)
    }
    return &Life{
        a: a, b: NewField(w, h),
        w: w, h: h,
        colorAlive: colorAlive, colorDead: colorDead,
    }
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
    // Update the state of the next field (b) from the current field (a).
    for y := 0; y < l.h; y++ {
        for x := 0; x < l.w; x++ {
            l.b.Set(x, y, l.a.Next(x, y))
        }
    }
    // Swap fields a and b.
    l.a, l.b = l.b, l.a
}

// String returns the game board as a string.
func (l *Life) String() string {
    var buf bytes.Buffer

    for y := 0; y < l.h; y++ {
        for x := 0; x < l.w; x++ {
            if l.a.Alive(x, y) {
                buf.WriteString(l.colorAlive)
                buf.WriteByte(' ')
            } else {
                buf.WriteString(l.colorDead)
                buf.WriteByte(' ')
            }

            buf.WriteString("\033[0m")
        }

        buf.WriteByte('\n')
    }

    return buf.String()
}
