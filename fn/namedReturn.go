package main
import "fmt"

func yearsUntilEvents(age int) ( yearsUntilAdult, yearsUntilCarRental, yearsUntilDrinking int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return 
}

func main() {
	adult, rental, drinking := yearsUntilEvents(16)
	fmt.Println("Years until adulthood:", adult)
	fmt.Println("Years until car rental:", rental)
	fmt.Println("Years until legal drinking:", drinking)
}
