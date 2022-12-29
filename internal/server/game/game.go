package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

type PlayerType int

const (
	HumanPlayer PlayerType = iota
	AIPlayer
)

type Player struct {
	Type PlayerType
}

type PlayerTurn int

const (
	Player1 PlayerTurn = iota
	Player2
)

type Game struct {
	GameMode   GameMode
	Board      Board
	PlayerTurn PlayerTurn
	Winner     Winner
	Players    [2]Player // 0 = player 1 (black), 1 = player 2 (white)
	TurnCount  int
}

type Move struct {
	Path             []int
	CheckersCaptured []int
}

func (g *Game) AIMove(playerTurn PlayerTurn) {
	// Get possible moves for every checker for g.Players[g.Playturn]
	possibleNonCapture := []Move{}
	possibleCapture := []Move{}

	loc := func(row, col int) int {
		return (row)*8 + (col)
	}

	for i, row := range g.Board.Positions {
		for j, square := range row {
			if playerTurn == Player1 && strings.ToLower(square) == "b" {
				n, c := g.Board.PossibleMoves(loc(i, j))
				possibleNonCapture = append(possibleNonCapture, n...)
				possibleCapture = append(possibleCapture, c...)
			} else if playerTurn == Player2 && strings.ToLower(square) == "w" {
				n, c := g.Board.PossibleMoves(loc(i, j))
				possibleNonCapture = append(possibleNonCapture, n...)
				possibleCapture = append(possibleCapture, c...)
			}
		}
	}

	rand.Seed(time.Now().UTC().Unix())

	// Pick a random possible move
	if len(possibleCapture) > 0 {
		r := rand.Intn(len(possibleCapture))
		g.Move(possibleCapture[r])
	} else if len(possibleNonCapture) > 0 {
		r := rand.Intn(len(possibleNonCapture))
		g.Move(possibleNonCapture[r])
	}
}

func (g *Game) Move(move Move) (followUpMoves []Move) {
	defer func() {
		g.TurnCount++
	}()

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

	// Promotion
	if to >= 56 && g.Board.Positions[to/8][to%8] == "b" {
		g.Board.Positions[to/8][to%8] = "B"
	} else if to <= 7 && g.Board.Positions[to/8][to%8] == "w" {
		g.Board.Positions[to/8][to%8] = "W"
	}

	if g.PlayerTurn == Player1 {
		g.PlayerTurn = Player2
	} else {
		g.PlayerTurn = Player1
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
	// p := g.Board.Positions
	// fmt.Printf("p: %+v", p)
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
