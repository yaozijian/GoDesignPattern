package m

import "fmt"

type (
	haierFactory        struct{}
	haierTV             struct{}
	haierWashingMachine struct{}
	haierAirConditioner struct{}
)

func NewHaierFactory() Factory {
	return &haierFactory{}
}

func (f *haierFactory) CreateTV() TV {
	return new(haierTV)
}

func (f *haierFactory) CreateWashingMachine() WashingMachine {
	return new(haierWashingMachine)
}

func (f *haierFactory) CreateAirConditioner() AirConditioner {
	return new(haierAirConditioner)
}

func (h *haierTV) Open() {
	fmt.Println("Haier TV Start...")
}

func (h *haierTV) Close() {
	fmt.Println("Haier TV Stop...")
}

func (h *haierWashingMachine) Start() {
	fmt.Println("Haier WashingMachine Start...")
}

func (h *haierWashingMachine) Stop() {
	fmt.Println("Haier WashingMachine Stop...")
}

func (h *haierAirConditioner) Start() {
	fmt.Println("Haier AirConditioner Start...")
}

func (h *haierAirConditioner) Stop() {
	fmt.Println("Haier AirConditioner Stop...")
}
