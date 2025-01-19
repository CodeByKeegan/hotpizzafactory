package api

type Message struct {
	PlayerId string

	Timestamp int

	Payload Payload
}

type Payload struct {
	Action GameAction
}

type GameAction int

const (
	Join GameAction = iota
	Leave
	Start
	End
)

type PlayerRole int

const (
	Host PlayerRole = iota
	Participant
)

type Player struct {
	id string

	name string

	role PlayerRole
}

type Game struct {
	players map[string]Player

	activePlayer Player

	state GameState
}

type GameState int

const (
	Waiting GameState = iota
	Playing
	Ended
)
