package api

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func MessageFromBytes(b []byte) Message {
	var message Message
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&message)
	if err != nil {
		panic(err)
	}
	return message
}

func BytesFromMessage(m Message) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

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
