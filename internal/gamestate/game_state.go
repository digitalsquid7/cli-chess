package gamestate

type GameState struct {
	Board *Board
}

func NewGameState() *GameState {
	return &GameState{
		Board: NewBoard(),
	}
}
