package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	// Cols represents number of columns.
	Cols = 9
	// Rows represents number of rows.
	Rows = 9
	// BoxSize represents size of square.
	BoxSize = 3
	// MaxValue represents max value of element on the board.
	MaxValue = 9
)

const (
	// Sud1 sudoku example
	Sud1 = `4 . . |. . . |8 . 5
            . 3 . |. . . |. . .
            . . . |7 . . |. . .
            ------+------+------
            . 2 . |. . . |. 6 .
            . . . |. 8 . |4 . .
            . . . |. 1 . |. . .
            ------+------+------
            . . . |6 . 3 |. 7 .
            5 . . |2 . . |. . .
            1 . 4 |. . . |. . .`

	// Sud2 sudoku example
	Sud2 = `. . . |. . . |6 8 .
            . . . |. 7 3 |. . 9
            3 . 9 |. . . |. 4 5
            ------+------+------
            4 9 . |. . . |. . .
            8 . 3 |. 5 . |9 . 2
            . . . |. . . |. 3 6
            ------+------+------
            9 6 . |. . . |3 . 8
            7 . . |6 8 . |. . .
            . 2 8 |. . . |. . .`
)

// Sudoku contains sudoku board
type Sudoku struct {
	s [][]int
}

func main() {
	s := Sudoku{}
	s.load(Sud1)
	s.solve(0, 0)
	s.printResult()
}

func (s *Sudoku) solve(r, c int) bool {
	if r > Rows-1 {
		return true
	}

	if s.s[r][c] != 0 {
		return s.nextField(r, c)
	}

	for v := 1; v <= MaxValue; v++ {
		if s.validVertically(c, v) && s.validHorizontally(r, v) && s.validBox(c, r, v) {
			s.s[r][c] = v
			res := s.nextField(r, c)
			if res {
				return true
			}
		}
	}

	s.s[r][c] = 0
	return false
}

func (s *Sudoku) nextField(r, c int) bool {
	if c >= Cols-1 { // next row
		return s.solve(r+1, 0)
	}
	// same row
	return s.solve(r, c+1)
}

func (s *Sudoku) printResult() {
	for _, val := range s.s {
		for _, val2 := range val {
			fmt.Printf("%d ", val2)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (s *Sudoku) validBox(c, r, v int) bool {
	row := (r / BoxSize) * BoxSize
	col := (c / BoxSize) * BoxSize

	for r := 0; r < BoxSize; r++ {
		for c := 0; c < BoxSize; c++ {
			if s.s[row+r][col+c] == v {
				return false
			}
		}
	}

	return true
}

func (s *Sudoku) validHorizontally(r, v int) bool {
	for _, elem := range s.s[r] {
		if elem == v {
			return false
		}
	}

	return true
}

func (s *Sudoku) validVertically(c, v int) bool {
	for i := 0; i < Rows; i++ {
		if s.s[i][c] == v {
			return false
		}
	}

	return true
}

func (s *Sudoku) load(su string) {
	s.s = make([][]int, Cols)
	lines := strings.Split(su, "\n")

	if len(lines) != Rows+2 {
		panic("Wrong format of sudoku")
	}

	offset := 0
	for i, l := range lines {
		str := strings.Trim(l, " ")

		if i == 3 || i == 7 {
			offset++
			continue // in those lines there are selectors only
		}

		chars := strings.FieldsFunc(str, filterValues)
		if len(chars) != 9 {
			panic("Wrong format of sudoku")
		}

		d := make([]int, Cols)
		for j, v := range chars {
			if v == "." {
				d[j] = int(0)
			} else {
				val, e := strconv.ParseUint(v, 10, 8)
				if e != nil {
					panic("Wrong format of sudoku. Invalid number")
				}
				d[j] = int(val)
			}
		}

		s.s[i-offset] = d
	}
}

func filterValues(c rune) bool {
	return !unicode.IsDigit(c) && c != '.'
}
