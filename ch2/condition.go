package main

import "fmt"

func main() {
	messageLe := 10
	maxMessageLe := 20
	fmt.Println("Trying to send a message of length:", messageLe, "and a max length of:", maxMessageLe)

	// don't touch above this line
	if messageLe <= maxMessageLe{
		fmt.Println("message sent")
	}else{
		fmt.Println("not send")
	}

}
