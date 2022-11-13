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
	return &Game{
		Board: Board{
			State: "_b_b_b_bb_b_b_b__b_b_b_b________________w_w_w_w__w_w_w_ww_w_w_w_",
		},
	}
}

func init() {
	GameState = NewGame()
}
