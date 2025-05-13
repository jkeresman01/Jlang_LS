package main

import (
	"fmt"
)

type Message struct {
	Rpc string `json:"rpc"`
	Id  int32  `json:"id"`
}

func CreateMessage(rpc string, id int32) *Message {
	return &Message{
		Rpc: rpc,
		Id:  id,
	}
}

func (message *Message) ToString() string {
	return fmt.Sprintf("id :%d rpc version: %s ", message.Id, message.Rpc)
}

func main() {
	message := CreateMessage("2.0", 1)
	fmt.Println(message.ToString())
}
