package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	var messages [3]string
	var costs [3]int

	messages[0] = primary
	messages[1] = secondary
	messages[2] = tertiary

	// Calculate cumulative lengths
	costs[0] = len(messages[0])
	costs[1] = costs[0] + len(messages[1])
	costs[2] = costs[1] + len(messages[2])

	return messages, costs
}
msgs, costs := getMessageWithRetries("hello", "world", "!")
