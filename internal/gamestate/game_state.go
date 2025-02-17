package gamestate

type GameState struct {
	Board  Board
	Cursor [2]int
}

func NewGameState() *GameState {
	return &GameState{
		Board:  NewBoard(),
		Cursor: [2]int{0, 2},
	}
}
