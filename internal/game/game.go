package game

import (
	"fmt"
	"strings"
)

type Winner int

const (
	NoWinner Winner = iota
	BlackWinner
	WhiteWinner
)

type GameMode int

const (
	NewGameMode GameMode = iota
	SinglePlayer
	TwoPlayer
)

type Game struct {
	GameMode   GameMode
	Board      Board
	PlayerTurn bool // false = black's turn; true = white's turn
	Winner     Winner
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

		_, captureMoves := g.Board.PossibleMoves(to)
		if len(captureMoves) > 0 {
			return captureMoves
		}
	}

	if to >= 56 && g.Board.Positions[to/8][to%8] == "b" {
		g.Board.Positions[to/8][to%8] = "B"
	} else if to <= 7 && g.Board.Positions[to/8][to%8] == "w" {
		g.Board.Positions[to/8][to%8] = "W"
	}

	if g.PlayerTurn {
		g.PlayerTurn = false
	} else {
		g.PlayerTurn = true
	}

	state := g.StateToString()
	if strings.ToLower(g.Board.Positions[to/8][to%8]) == "b" && !strings.Contains(state, "w") && !strings.Contains(state, "W") {
		g.Winner = BlackWinner
	} else if strings.ToLower(g.Board.Positions[to/8][to%8]) == "w" && !strings.Contains(state, "b") && !strings.Contains(state, "B") {
		g.Winner = WhiteWinner
	}

	return
}

// func (b *Board) isKingPromotion(checkerLocation, r, c int) bool {
// 	if b.state checkerLocation b.Positions[r][c] == "b"
// }

func (g *Game) StateToString() string {
	p := g.Board.Positions
	fmt.Printf("p: %+v", p)
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
