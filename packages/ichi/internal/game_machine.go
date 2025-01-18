package api

import "fmt"

func NewGame() *Game {
	return &Game{
		players: []Player{},
		state:   Waiting,
	}
}

func NextAction(game *Game, message Message) Message {
	action := message.Payload.Action
	switch action {
	case Join:
		// Do something
		fmt.Println("Join", message.PlayerId)
	case Leave:
		// Do something
		fmt.Println("Leave", message.PlayerId)
	case Start:
		// Do something
		fmt.Println("Start", message.PlayerId)
	case End:
		// Do something
		fmt.Println("End", message.PlayerId)
	default:
		panic(fmt.Errorf("unknown state: %d", action))
	}

	return message
}


