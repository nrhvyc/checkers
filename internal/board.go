package internal

// Board - checkers game state
type Board struct{}

func newBoard() (board Board) {
	return Board{}
}

// func (b *Board) String() {
// 	for row := range b {
// 		for square := range {
// 			fmt.Println(square)
// 			s := strings.Replace(row, "", "-")
// 			fmt.Println(s)
// 		}
// 	}
// }
