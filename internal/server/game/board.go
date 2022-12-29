package game

import (
	"fmt"
	"strings"
)

type Board struct {
	state string

	Positions [8][]string // 8 x 8 grid
	// BlackCheckerPositions map[string]int // map[color]location
	// WhiteCheckerPositions map[string]int // map[color]location
}

func (b *Board) PossibleMoves(checkerLocation int) (nonCaptureMoves, captureMoves []Move) {
	r := checkerLocation / 8
	c := checkerLocation % 8
	loc := func(r, c int) int {
		return (r)*8 + (c)
	}

	if strings.ToLower(b.Positions[r][c]) == "b" {
		// Black Moves
		if r+1 < 8 || b.Positions[r][c] == "B" {
			if b.isEmptyAndValid(r+1, c-1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r+1, c-1)},
				})
			}
			if b.isEmptyAndValid(r+1, c+1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r+1, c+1)},
				})
			}
			if b.containsWhiteAndValid(r+1, c-1) && b.isEmptyAndValid(r+2, c-2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r+2, c-2)},
					CheckersCaptured: []int{loc(r+1, c-1)},
				})
			}
			if b.containsWhiteAndValid(r+1, c+1) && b.isEmptyAndValid(r+2, c+2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r+2, c+2)},
					CheckersCaptured: []int{loc(r+1, c+1)},
				})
			}
		}
		if b.Positions[r][c] == "B" {
			if c-1 > 0 && b.isEmptyAndValid(r-1, c-1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r-1, c-1)},
				})
			}
			if c+1 < 8 && b.isEmptyAndValid(r-1, c+1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r-1, c+1)},
				})
			}
			if b.containsWhiteAndValid(r-1, c-1) && b.isEmptyAndValid(r-2, c-2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r-2, c-2)},
					CheckersCaptured: []int{loc(r-1, c-1)},
				})
			}
			if b.containsWhiteAndValid(r-1, c+1) && b.isEmptyAndValid(r-2, c+2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r-2, c+2)},
					CheckersCaptured: []int{loc(r-1, c+1)},
				})
			}
		}
	} else {
		// White Moves
		if r-1 >= 0 || b.Positions[r][c] == "W" {
			if c-1 > 0 && b.isEmptyAndValid(r-1, c-1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r-1, c-1)},
				})
			}
			if c+1 < 8 && b.isEmptyAndValid(r-1, c+1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r-1, c+1)},
				})
			}
			if b.containsBlackAndValid(r-1, c-1) && b.isEmptyAndValid(r-2, c-2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r-2, c-2)},
					CheckersCaptured: []int{loc(r-1, c-1)},
				})
			}
			if b.containsBlackAndValid(r-1, c+1) && b.isEmptyAndValid(r-2, c+2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r-2, c+2)},
					CheckersCaptured: []int{loc(r-1, c+1)},
				})
			}
		}
		if b.Positions[r][c] == "W" {
			if b.isEmptyAndValid(r+1, c-1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r+1, c-1)},
				})
			}
			if b.isEmptyAndValid(r+1, c+1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r+1, c+1)},
				})
			}
			if b.containsBlackAndValid(r+1, c-1) && b.isEmptyAndValid(r+2, c-2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r+2, c-2)},
					CheckersCaptured: []int{loc(r+1, c-1)},
				})
			}
			if b.containsBlackAndValid(r+1, c+1) && b.isEmptyAndValid(r+2, c+2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r+2, c+2)},
					CheckersCaptured: []int{loc(r+1, c+1)},
				})
			}
		}
	}

	fmt.Printf("PossibleMoves() nonCaptureMoves: %+v captureMoves: %+v\n", nonCaptureMoves, captureMoves)

	return
}

func (b *Board) isEmptyAndValid(r, c int) bool {
	if r > 7 || r < 0 || c > 7 || c < 0 {
		return false
	}
	return b.Positions[r][c] == "_"
}

func (b *Board) containsWhiteAndValid(r, c int) bool {
	if r > 7 || r < 0 || c > 7 || c < 0 {
		return false
	}
	return strings.ToLower(b.Positions[r][c]) == "w"
}

func (b *Board) containsBlackAndValid(r, c int) bool {
	if r > 7 || r < 0 || c > 7 || c < 0 {
		return false
	}
	return strings.ToLower(b.Positions[r][c]) == "b"
}
