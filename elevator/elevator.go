package elevator

import "fmt"

type ElevatorBank struct {
	Total     int        `json:"total"`
	Elevators []Elevator `json:"elevators"`
}

type Elevator struct {
	Destinations []int  `json:"destinations"`
	Status       string `json:"status"`
	CurrentFloor int    `json:"current_floor"`
}

func NewElevatorBank(total int) *ElevatorBank {
	elevators := make([]Elevator, 0)

	for i := 0; i < total; i++ {
		elevators = append(elevators, Elevator{Status: "idle"})
	}

	return &ElevatorBank{
		Total:     total,
		Elevators: elevators,
	}
}

func (eb *ElevatorBank) Status() string {
	return fmt.Sprintf("Number of elevators: %d", eb.Total)
}

func (eb *ElevatorBank) SaveToDisk() error {
	pers := NewPersistence(eb, "elevators.json")
	return pers.Save()
}
