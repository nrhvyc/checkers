package game

var GameState *Game

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
