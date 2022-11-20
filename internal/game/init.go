package game

var GameState *Game

// _b_b_b_b
// b_b_b_b_
// ___b_b_b
// b_______
// ________
// w_w_w_w_
// _w_w_w_w
// w_w_w_w_
func NewGame() *Game {
	g := Game{
		Board: Board{
			// state: "_b_b_b_bb_b_b_b__b_b_b_b________________w_w_w_w__w_w_w_ww_w_w_w_",
			// state: "_b_b_b_bb_b_b_b____b_b________b__b___w__w_w___w__w_w_w_ww_w_w_w_", // Test Case for multi capture
			// state: "_b_b___bb_b_b_b____b_b____b___b__b_w_w__w___w_w__w___w_ww_w_w_w_", // Test Case for multi capture 2
			// state: "_b_b_w_bb___b____b_b______b_b____b_w____w_w_w_w______w_ww_w_w_w_", // Test Case for multi capture 3
			// state: "_b___w_bb_b_b____b_b______b_b____b_w___ww_w_w________w_ww_w_w_w_",  // Test Case for multi capture 4
			state: " ____________________________b______w____________________________", // Test Case end game
		},
	}
	g.SetStateFromString()
	return &g
}

func init() {
	GameState = NewGame()
}
