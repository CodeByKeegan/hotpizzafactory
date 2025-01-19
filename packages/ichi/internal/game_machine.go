package api

import "fmt"

func NewGame() *Game {
	return &Game{
		players: make(map[string]Player),
		state:   Waiting,
	}
}

func NewGameStatus(game *Game) GameStatus{
	// player hands
	playerHands := make(map[int]string)
	for _, player := range game.players {
		playerHands[player.hand] = player.id
	}

	// player host
	playerHost := ""
	for _, player := range game.players {
		if player.role == Host {
			playerHost = player.id
			break
		}
	}

	return GameStatus{
		PlayerHands: playerHands,
		Host:  playerHost,
		ActivePlayerId: game.activePlayer.id,
		State:	game.state,
	}
}

func NextAction(game *Game, message Message) GameStatus {
	msgPlayer, ok := game.players[message.PlayerId]

	action := message.Payload.Action
	switch action {
	case Join:
		// if msgPlayer is already in the game, do nothing
		// if msgPlayer is not in the game, create a new player & add to the game
		if !ok {
			// create a new player
			msgPlayer = Player{ id: message.PlayerId, name: "Player" }

			// determine player role
			if game.state != Waiting {
				msgPlayer.role = Audience
			} else {
				if len(game.players) == 0 {
					msgPlayer.role = Host
				} else {
					msgPlayer.role = Participant
				}
			}

			// add player to the game
			game.players[message.PlayerId] = msgPlayer
		}
		
		fmt.Println("Join", message.PlayerId, msgPlayer.role, len(game.players))
	case Leave:
		// if msgPlayer is not in the game, do nothing
		// if msgPlayer is the host, assign new host and remove
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

		fmt.Println("Leave", message.PlayerId, len(game.players))
	case Start:
		// if msgPlayer is not the host, do nothing
		// if msgPlayer is the host, change game state to Playing
		if msgPlayer.role == Host {
			game.state = Playing
		}
		fmt.Println("Start", message.PlayerId)
	case End:
		// if msgPlayer is not the host, do nothing
		// if msgPlayer is the host, change game state to Ended
		if msgPlayer.role == Host {
			game.state = Ended
		}
		fmt.Println("End", message.PlayerId)

	default:
		panic(fmt.Errorf("unknown state: %d", action))
	}

	return NewGameStatus(game)
}


