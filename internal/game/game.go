package game

import (
	"fmt"
	"strings"
)

type Game struct {
	Board      Board
	PlayerTurn bool // false = black's turn; true = white's turn
}

type Board struct {
	state string

	Positions [8][]string // 8 x 8 grid
	// BlackCheckerPositions map[string]int // map[color]location
	// WhiteCheckerPositions map[string]int // map[color]location
}

type Move struct {
	Path             []int
	CheckersCaptured []int
}

func (g *Game) Move(move Move) (followUpMoves []Move) {
	from := move.Path[0]
	to := move.Path[len(move.Path)-1]
	fmt.Printf("game.Move() from: %d to: %d", from, to)

	g.Board.Positions[to/8][to%8] = g.Board.Positions[from/8][from%8]
	g.Board.Positions[from/8][from%8] = "_"

	// remove most recent checker capture
	if lenCheckers := len(move.CheckersCaptured); lenCheckers > 0 {
		fmt.Printf("captureLocation: %d", move.CheckersCaptured[lenCheckers-1])
		g.Board.Positions[move.CheckersCaptured[lenCheckers-1]/8][move.CheckersCaptured[lenCheckers-1]%8] = "_"

		_, captureMoves := g.PossibleMoves(to)
		if len(captureMoves) > 0 {
			return captureMoves
		}

	}

	if g.PlayerTurn {
		g.PlayerTurn = false
	} else {
		g.PlayerTurn = true
	}
	return
}

func (g *Game) StateToString() string {
	state := make([]string, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			state[i*8+j] = g.Board.Positions[i][j]
		}
	}
	out := strings.Join(state, "")
	return out
}

func (g *Game) SetStateFromString() {
	p := strings.Split(g.Board.state, "")
	for i := 0; i < 8; i++ {
		g.Board.Positions[i] = make([]string, 8)
		for j := 0; j < 8; j++ {
			g.Board.Positions[i][j] = p[i*8+j]
		}
	}
	fmt.Printf("g.Board.Positions: %+v\n", g.Board.Positions)
}

func (g *Game) PossibleMoves(checkerLocation int) (nonCaptureMoves, captureMoves []Move) {
	r := checkerLocation / 8
	c := checkerLocation % 8
	loc := func(r, c int) int {
		return (r)*8 + (c)
	}

	if g.Board.Positions[r][c] == "b" {
		// Black Moves
		if r+1 < 8 {
			if g.Board.isEmptyAndValid(r+1, c-1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r+1, c-1)},
				})
			}
			if g.Board.isEmptyAndValid(r+1, c+1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r+1, c+1)},
				})
			}
			if g.Board.containsWhiteAndValid(r+1, c-1) && g.Board.isEmptyAndValid(r+2, c-2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r+2, c-2)},
					CheckersCaptured: []int{loc(r+1, c-1)},
				})
			}
			if g.Board.containsWhiteAndValid(r+1, c+1) && g.Board.isEmptyAndValid(r+2, c+2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r+2, c+2)},
					CheckersCaptured: []int{loc(r+1, c+1)},
				})
			}
		}
	} else {
		// White Moves
		if r-1 >= 0 {
			if c-1 > 0 && g.Board.isEmptyAndValid(r-1, c-1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r-1, c-1)},
				})
			}
			if c+1 < 8 && g.Board.isEmptyAndValid(r-1, c+1) {
				nonCaptureMoves = append(nonCaptureMoves, Move{
					Path: []int{loc(r, c), loc(r-1, c+1)},
				})
			}
			if g.Board.containsBlackAndValid(r-1, c-1) && g.Board.isEmptyAndValid(r-2, c-2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r-2, c-2)},
					CheckersCaptured: []int{loc(r-1, c-1)},
				})
			}
			if g.Board.containsBlackAndValid(r-1, c+1) && g.Board.isEmptyAndValid(r-2, c+2) {
				captureMoves = append(captureMoves, Move{
					Path:             []int{loc(r, c), loc(r-2, c+2)},
					CheckersCaptured: []int{loc(r-1, c+1)},
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
	return b.Positions[r][c] == "w"
}

func (b *Board) containsBlackAndValid(r, c int) bool {
	if r > 7 || r < 0 || c > 7 || c < 0 {
		return false
	}
	return b.Positions[r][c] == "b"
}
