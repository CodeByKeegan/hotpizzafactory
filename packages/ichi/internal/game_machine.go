package api

import "fmt"

func NewGame() *Game {
	return &Game{
		players: make(map[string]Player),
		state:   Waiting,
	}
}

func NextAction(game *Game, message Message) Message {
	msgPlayer, ok := game.players[message.PlayerId]

	action := message.Payload.Action
	switch action {
	case Join:
		// Do something
		if !ok {
			msgPlayer = Player{ id: message.PlayerId, role: Participant }
			game.players[message.PlayerId] = msgPlayer
		}
		
		fmt.Println("Join", message.PlayerId)
	case Leave:
		// if msgPlayer is not in the game, do nothing
		// if msgPlayer is the host, assign new host
		// if msgPlayer is the participant, remove from the game
		if ok {
			if msgPlayer.role == Host {
				// assign new host
				for _, player := range game.players {
					if player.role == Participant {
						player.role = Host
						break
					}
				}
			}
			delete(game.players, message.PlayerId)
		}

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

	println(len(game.players))

	// TODO: Send message version of game state to players
	return message
}


