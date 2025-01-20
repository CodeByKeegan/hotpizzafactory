package api

import "log"

type Message struct {
	PlayerId string

	Timestamp int

	Event Event
}

type Event interface {
	IsEvent()
}

type EmptyEvent struct {
	Action GameAction
}

type CardEvent struct {
	Action GameAction
	PlayedCard Card
}

func (EmptyEvent) IsEvent() {}
func (CardEvent) IsEvent() {}

func HandleEvent(ev Event) {
	switch evt := ev.(type) {
	case CardEvent:
		log.Println("PlayCardEvent", evt.PlayedCard)
	default:
		log.Println("Default")
	}
}


type GameAction int

const (
	Join GameAction = iota
	Leave
	Start
	End
	PlayCard
	DrawCard
)

type PlayerRole int

const (
	Host PlayerRole = iota
	Participant
	Audience
)

type User struct {
	id string

	name string

	role PlayerRole

	hand []Card
}

type Player struct {
	Id string

	Name string

	Hand int
}

type Game struct {
	players map[string]User

	playerOrder []string

	activePlayer User

	state GameState

	deck []Card

	discard []Card
}

type GameState int

const (
	Waiting GameState = iota
	Playing
	Ended
)

type ClientState struct {
	Players []Player

	Order []string

	Host string

	ActivePlayerId string

	State GameState

	LastPlayedCard Card

	Hand []Card
}
