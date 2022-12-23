package game

import (
	"math/rand"
	"time"
)

var GameState *Game

func NewGame(gameMode GameMode) *Game {
	// Set Play Positions
	var players [2]Player
	if gameMode == SinglePlayer {
		rand.Seed(time.Now().UTC().Unix())
		pos := rand.Intn(2)
		players[pos] = Player{
			Type: HumanPlayer,
		}
		aiPos := 0
		if pos == 0 {
			aiPos = 1
		}
		players[aiPos] = Player{
			Type: AIPlayer,
		}
		// client.NewPeer()
	}

	g := Game{
		GameMode: gameMode,
		Board: Board{
			state: "_b_b_b_bb_b_b_b__b_b_b_b________________w_w_w_w__w_w_w_ww_w_w_w_",
			// state: "_b_b_b_bb_b_b_b____b_b________b__b___w__w_w___w__w_w_w_ww_w_w_w_", // Test Case for multi capture
			// state: "_b_b___bb_b_b_b____b_b____b___b__b_w_w__w___w_w__w___w_ww_w_w_w_", // Test Case for multi capture 2
			// state: "_b_b_w_bb___b____b_b______b_b____b_w____w_w_w_w______w_ww_w_w_w_", // Test Case for multi capture 3
			// state: "_b___w_bb_b_b____b_b______b_b____b_w___ww_w_w________w_ww_w_w_w_",  // Test Case for multi capture 4
			// state: " ____________________________b______w____________________________", // Test Case end game
		},
		Players:   players,
		TurnCount: 0,
	}
	g.SetStateFromString()
	return &g
}

// func init() {
// 	GameState = NewGame()
// }
