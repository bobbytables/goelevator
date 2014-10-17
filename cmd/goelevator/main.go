package main

import (
	"github.com/bobbytables/goelevator/elevator"
)

func main() {
	eb := elevator.NewElevatorBank(3)

	eb.SaveToDisk()

	println(eb.Status())
}
