package api

import (
	"encoding/json"
)

func MessageFromJson(b []byte) Message {
	var message Message
	err := json.Unmarshal(b, &message)
	if err != nil {
		panic(err)
	}
	return message
}

// JsonFromMessage encodes a Message to a JSON byte array
func JsonFromMessage(m Message) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}

// JsonFromClientState encodes a ClientState to a JSON byte array
func JsonFromClientState(gs ClientState) []byte {
	b, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return b
}
