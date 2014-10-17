package elevator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ElevatorBankPersistence struct {
	ElevatorBank *ElevatorBank
	Filename     string
}

func NewPersistence(eb *ElevatorBank, filename string) *ElevatorBankPersistence {
	return &ElevatorBankPersistence{
		ElevatorBank: eb,
		Filename:     filename,
	}
}

func (p *ElevatorBankPersistence) Contents() ([]byte, error) {
	return ioutil.ReadFile(p.Filename)
}

func (p *ElevatorBankPersistence) Save() error {
	contents, err := json.Marshal(p.ElevatorBank)

	if err != nil {
		panic(fmt.Sprintf("Could not save file: %s", err))
	}

	return ioutil.WriteFile(p.Filename, contents, 0777)
}

func (p *ElevatorBankPersistence) BuildElevatorBank() *ElevatorBank {
	bank := &ElevatorBank{}
	contents, err := p.Contents()

	if err != nil {
		panic(fmt.Sprintf("Could not load file: %s", err))
	}

	json.Unmarshal(contents, bank)

	return bank
}
