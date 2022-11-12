package game

type Game struct {
	Board Board `json:"board"`
}

type Board struct {
	State string `json:"state"`
}
