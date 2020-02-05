package main

// Field represents a two-dimensional field of cells.
type Field struct {
    s    [][]bool
    w, h int
}

// NewField returns an empty field of the specified width and height.
func NewField(w, h int) *Field {
    s := make([][]bool, h)
    for i := range s {
        s[i] = make([]bool, w)
    }
    return &Field{s: s, w: w, h: h}
}

// Set sets the state of the specified cell to the given value.
func (f *Field) Set(x, y int, b bool) {
    f.s[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the x or y coordinates are outside the field boundaries they are wrapped
// toroidally. For instance, an x value of -1 is treated as width-1.
func (f *Field) Alive(x, y int) bool {
    x += f.w
    x %= f.w
    y += f.h
    y %= f.h
    return f.s[y][x]
}

// Next returns the state of the specified cell at the next time step.
func (f *Field) Next(x, y int) bool {
    // Count the adjacent cells that are alive.
    alive := 0
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
                alive++
            }
        }
    }
    // Return next state according to the game rules:
    //   exactly 3 neighbors: on,
    //   exactly 2 neighbors: maintain current state,
    //   otherwise: off.
    return alive == 3 || (alive == 2 && f.Alive(x, y))
}
