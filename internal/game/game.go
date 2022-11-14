package game

import (
	"fmt"
	"strings"
)

type Game struct {
	Board Board
}

type Board struct {
	state string

	Positions [8][]string // 8 x 8 grid
	// BlackCheckerPositions map[string]int // map[color]location
	// WhiteCheckerPositions map[string]int // map[color]location
}

func (g *Game) Move(from, to int) {
	g.Board.Positions[to/8][to%8] = g.Board.Positions[from/8][from%8]
	g.Board.Positions[from/8][from%8] = "_"
}

func (g *Game) StateToString() string {
	state := make([]string, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			state[i*8+j] = g.Board.Positions[i][j]
		}
	}
	out := strings.Join(state, "")
	fmt.Printf("out: %s\n", out)
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

func (g *Game) PossibleMoves(checkerLocation int) []int {
	i := checkerLocation / 8
	j := checkerLocation % 8

	locations := []int{}

	if g.Board.Positions[i][j] == "b" {
		// Black Moves
		if i+1 < 8 {
			if j-1 >= 0 && g.Board.isEmpty(i+1, j-1) {
				locations = append(locations, (i+1)*8+(j-1))
			}
			if j+1 < 8 && g.Board.isEmpty(i+1, j+1) {
				locations = append(locations, (i+1)*8+(j+1))
			}
		}
	} else {
		// White Moves
		if i-1 >= 0 {
			if j-1 >= 0 && g.Board.isEmpty(i-1, j-1) {
				locations = append(locations, (i-1)*8+(j-1))
			}
			if j+1 < 8 && g.Board.isEmpty(i-1, j+1) {
				locations = append(locations, (i-1)*8+(j+1))
			}
		}
	}

	return locations
}

func (b *Board) isEmpty(i, j int) bool {
	return b.Positions[i][j] == "_"
}
