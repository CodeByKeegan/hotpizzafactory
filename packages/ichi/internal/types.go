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
	Play
)

type PlayerRole int

const (
	Host PlayerRole = iota
	Participant
	Audience
)

type Player struct {
	id string

	name string

	role PlayerRole

	hand int
}

type Game struct {
	players map[string]Player

	playerOrder []string

	activePlayer Player

	state GameState
}

type GameState int

const (
	Waiting GameState = iota
	Playing
	Ended
)

type ClientState struct {
	PlayerHands map[string]int

	PlayerOrder []string

	Host string

	ActivePlayerId string 

	State GameState
}
