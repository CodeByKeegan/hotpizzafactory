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
			go func() {
				h.broadcast <- JsonFromMessage(Message{PlayerId: client.playerId, Event: EventUnion{ Join: &JoinEvent{} }})
			}()
		case client := <-h.unregister:
			println("Unregister")
			go func() {
				h.broadcast <- JsonFromMessage(Message{PlayerId: client.playerId, Event: EventUnion{ Leave: &LeaveEvent{} }})
			}()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			println("Broadcast")
			HandlePlayerAction(h.game, MessageFromJson(message))
			for client := range h.clients {
				select {
				case client.send <- JsonFromClientState(NewClientStatus(h.game, client.playerId)):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
