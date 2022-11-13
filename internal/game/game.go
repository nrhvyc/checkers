package game

import "strings"

type Game struct {
	Board Board `json:"board"`
}

type Board struct {
	State string `json:"state"`
}

func (g *Game) Move(from, to int) {
	positions := strings.Split(g.Board.State, "")
	positions[to] = positions[from]
	positions[from] = "_"
	g.Board.State = strings.Join(positions, "")
}
