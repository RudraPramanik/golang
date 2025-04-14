package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.5
	msg:= fmt.Sprintf("hi %s, your open rate is OPENRATE %.2f percent \n", name, openRate)
	// ?

	// don't edit below this line

	fmt.Print(msg)
}
