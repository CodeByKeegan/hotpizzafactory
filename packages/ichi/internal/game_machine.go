package api

import (
	"fmt"
)

const (
	MinPlayerCount = 2
)

func NewGame() *Game {
	deck, discard := Draw(NewDeck(), 1)
	return &Game{
		players:     make(map[string]User),
		state:       Waiting,
		playerOrder: make([]string, 0),
		deck:        deck,
		discard:     discard,
	}
}

func NewClientStatus(game *Game, playerId string) ClientState {
	// players
	players := make([]Player, 0, len(game.players))
	for _, player := range game.players {
		players = append(players, Player{
			Id:   player.id,
			Name: player.name,
			Hand: len(player.hand),
		})
	}

	// player host
	playerHost := ""
	for _, player := range game.players {
		if player.role == Host {
			playerHost = player.id
			break
		}
	}

	var lastPlayedCard Card
	var currentPlayerHand []Card
	if game.state == Playing {
		lastPlayedCard = game.discard[len(game.discard)-1]
		currentPlayerHand = game.players[playerId].hand
	}

	return ClientState{
		Players:        players,
		Order:          game.playerOrder,
		Host:           playerHost,
		ActivePlayerId: game.activePlayer.id,
		State:          game.state,
		LastPlayedCard: lastPlayedCard,
		Hand:           currentPlayerHand,
	}
}

func HandlePlayerAction(game *Game, message Message) {
	player, ok := game.players[message.PlayerId]

	switch {
	case message.Event.Join != nil:
		// if msgPlayer is already in the game, do nothing
		// if msgPlayer is not in the game, create a new player & add to the game
		if !ok {
			// create a new player
			player = User{id: message.PlayerId, name: "Player"}

			// determine player role
			if game.state != Waiting {
				player.role = Audience
			} else if len(game.players) == 0 {
				player.role = Host
			} else {
				player.role = Participant
			}

			// add player to the game
			game.players[message.PlayerId] = player
		}

		fmt.Printf("Action:\t\t%s\nPlayer Id:\t%s\nPlayer Role:\t%s\nPlayer Count:\t%d\n", message.Event.Join.Action(), message.PlayerId, player.role, len(game.players))

	case message.Event.Leave != nil:
		// if msgPlayer is not in the game, do nothing
		// if msgPlayer is the host, assign new host and remove
		// if msgPlayer is the participant, remove from the game
		if ok {
			RemovePlayer(game, message.PlayerId)
		}

		fmt.Printf("Action:\t\t%s\nPlayer Id:\t%s\nPlayer Role:\t%s\nPlayer Count:\t%d\n", message.Event.Leave.Action(), message.PlayerId, player.role, len(game.players))

	case message.Event.Start != nil:
		// if msgPlayer is not the host, do nothing
		// if msgPlayer is the host, change game state to Playing
		if player.role == Host &&
			game.state == Waiting &&
			len(game.players) >= MinPlayerCount {
			StartGame(game)
		}
		fmt.Printf("Action:\t\t%s\nPlayer Id:\t%s\nPlayer Count:\t%d\n", message.Event.Start.Action(), message.PlayerId, len(game.players))

	case message.Event.End != nil:
		// if msgPlayer is not the host, do nothing
		// if msgPlayer is the host, change game state to Ended
		if player.role == Host &&
			game.state == Playing {
			game.state = Ended
		}
		fmt.Printf("Action:\t\t%s\nPlayer Id:\t%s\n", message.Event.End.Action(), message.PlayerId)

	case message.Event.PlayCard != nil:
		// if msgPlayer is not the active player, do nothing
		// if msgPlayer is the active player, decrement the hand count and change active player
		if player.id == game.activePlayer.id {

			// decrement hand count
			//msgPlayer.hand--
			//game.players[msgPlayer.id] = msgPlayer

			// if hand count is 0, end the game
			if len(player.hand) == 0 {
				game.state = Ended
				break
			}

			// change active player
			activePlayerIndex := 0
			for i, playerId := range game.playerOrder {
				if playerId == game.activePlayer.id {
					activePlayerIndex = (i + 1) % len(game.playerOrder)
					break
				}
			}

			game.activePlayer = game.players[game.playerOrder[activePlayerIndex]]

			fmt.Printf("Action:\t\t%s\nPlayer Id:\t%s\nCard Played:\t%s\n", message.Event.PlayCard.Action(), message.PlayerId, message.Event.PlayCard.Card)
		} else {
			fmt.Printf("Action:\t\t%s\nPlayer Id:\t%s\nCard Played:\tPlayed out of turn\n", message.Event.PlayCard.Action(), message.PlayerId)
		}

	default:
		panic(fmt.Errorf("invalid event"))
	}
}

func StartGame(game *Game) {
	game.state = Playing

	// set player order
	for playerId := range game.players {
		game.playerOrder = append(game.playerOrder, playerId)
	}

	// set active player
	game.activePlayer = game.players[game.playerOrder[0]]

	// create deck
	game.deck = NewDeck()

	// deal cards
	// for each player, draw 7 cards
	for _, player := range game.players {
		for i := 0; i < 7; i++ {
			game.deck, player.hand = Draw(game.deck, 7)
		}
		game.players[player.id] = player
	}
}

func RemovePlayer(game *Game, playerId string) {
	// if player is the host, assign new host
	// if player is the active player, change active player to next player
	// remove player from the player order
	// remove player from the game

	player := game.players[playerId]
	if player.role == Host {
		for _, newHost := range game.players {
			if newHost.role == Participant {
				newHost.role = Host
				game.players[newHost.id] = newHost
				break
			}
		}
	}

	if game.activePlayer.id == playerId {
		activePlayerIndex := 0
		for i, playerId := range game.playerOrder {
			if playerId == game.activePlayer.id {
				activePlayerIndex = i
				break
			}
		}

		activePlayerIndex = (activePlayerIndex + 1) % len(game.playerOrder)
		game.activePlayer = game.players[game.playerOrder[activePlayerIndex]]
	}

	playerOrder := make([]string, 0, len(game.playerOrder)-1)
	for _, playerId := range game.playerOrder {
		if playerId != player.id {
			playerOrder = append(playerOrder, playerId)
		}
	}
	game.playerOrder = playerOrder

	delete(game.players, playerId)
}
