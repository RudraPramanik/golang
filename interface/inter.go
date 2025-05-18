package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	// ?
	content := msg.getMessage()
	cost := len.(content)*3
	return content, cost
}

type message interface {
	// ?
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}


func main() {
	birthday := birthdayMessage{
		birthdayTime:  time.Date(2025, 4, 14, 0, 0, 0, 0, time.UTC),
		recipientName: "Alice",
	}
	content, cost := sendMessage(birthday)
	fmt.Println(content)
	fmt.Println("Cost:", cost)

	report := sendingReport{
		reportName:    "Weekly Summary",
		numberOfSends: 123,
	}
	content, cost = sendMessage(report)
	fmt.Println(content)
	fmt.Println("Cost:", cost)
}
