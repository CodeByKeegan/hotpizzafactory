package api

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Game
	game *Game
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		game:       NewGame(),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			println("Register")
			h.clients[client] = true
			message := JsonFromMessage(Message{PlayerId: client.playerId, Event: Event{Action: Join}})
			go func() {
				h.broadcast <- message
			}()
		case client := <-h.unregister:
			println("Unregister")
			message := JsonFromMessage(Message{PlayerId: client.playerId, Event: Event{Action: Leave}})
			go func() {
				h.broadcast <- message
			}()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			println("Broadcast")
			msg := MessageFromJson(message)
			NextAction(h.game, msg)
			for client := range h.clients {
				nextMsg := JsonFromClientState(NewClientStatus(h.game, client.playerId))
				select {
				case client.send <- nextMsg:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
