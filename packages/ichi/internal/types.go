package api

type Message struct {
	PlayerId string

	Timestamp int

	Event EventUnion
}

type EventUnion struct {
	Join *JoinEvent
	Leave *LeaveEvent
	Start *StartEvent
	End *EndEvent
	PlayCard *PlayCardEvent
	DrawCard *DrawCardEvent
}

type Event interface {
	Action() GameAction
}

type JoinEvent struct {}
func (JoinEvent) Action() GameAction { return Join }

type LeaveEvent struct {}
func (LeaveEvent) Action() GameAction { return Leave }

type StartEvent struct {}
func (StartEvent) Action() GameAction { return Start }

type EndEvent struct {}
func (EndEvent) Action() GameAction { return End }

type PlayCardEvent struct {
	Card Card
	TargetPlayer *string
}
func (PlayCardEvent) Action() GameAction { return PlayCard }

type DrawCardEvent struct {}
func (DrawCardEvent) Action() GameAction { return DrawCard }

type GameAction int

const (
	Join GameAction = iota
	Leave
	Start
	End
	PlayCard
	DrawCard
)

var gameActionName = map[GameAction]string{
	Join:     "Join",
	Leave:    "Leave",
	Start:    "Start",
	End:      "End",
	PlayCard: "PlayCard",
	DrawCard: "DrawCard",
}

func (ga GameAction) String() string {
	return gameActionName[ga]
}

type PlayerRole int

const (
	Host PlayerRole = iota
	Participant
	Audience
)

var playerRoleName = map[PlayerRole]string{
	Host:        "Host",
	Participant: "Participant",
	Audience:    "Audience",
}

func (pr PlayerRole) String() string {
	return playerRoleName[pr]
}

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

var gameStateName = map[GameState]string{
	Waiting: "Waiting",
	Playing: "Playing",
	Ended:   "Ended",
}

func (gs GameState) String() string {
	return gameStateName[gs]
}

type ClientState struct {
	Players []Player

	Order []string

	Host string

	ActivePlayerId string

	State GameState

	LastPlayedCard Card

	Hand []Card
}
