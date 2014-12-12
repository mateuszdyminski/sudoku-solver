package main

import "testing"

var result bool

func benchmarkSudoku(sud string, b *testing.B) {
	s := Sudoku{}
	for i := 0; i < b.N; i++ {
		s.load(sud)
		result = s.solve(0, 0)
	}
}

func BenchmarkSud1(b *testing.B) { benchmarkSudoku(Sud1, b) }
func BenchmarkSud2(b *testing.B) { benchmarkSudoku(Sud2, b) }
