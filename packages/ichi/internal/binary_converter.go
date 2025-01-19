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

// JsonFromGameStatus encodes a GameStatus to a JSON byte array
func JsonFromGameStatus(gs GameStatus) []byte {
	b, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return b
}
