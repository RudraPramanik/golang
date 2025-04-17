package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func main() {
	msg := messageToSend{
		message: "Hello!",
		sender: user{
			name:   "Alice",
			number: 123456,
		},
		recipient: user{
			name:   "Bob",
			number: 987654,
		},
	}

	fmt.Println(canSendMessage(msg))
}
