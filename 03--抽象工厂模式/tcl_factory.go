package m

import "fmt"

type (
	tclFactory        struct{}
	tclTV             struct{}
	tclWashingMachine struct{}
	tclAirConditioner struct{}
)

func NewTCLFactory() Factory {
	return &tclFactory{}
}

func (f *tclFactory) CreateTV() TV {
	return new(tclTV)
}

func (f *tclFactory) CreateWashingMachine() WashingMachine {
	return new(tclWashingMachine)
}

func (f *tclFactory) CreateAirConditioner() AirConditioner {
	return new(tclAirConditioner)
}

func (h *tclTV) Open() {
	fmt.Println("TCL TV Start...")
}

func (h *tclTV) Close() {
	fmt.Println("TCL TV Stop...")
}

func (h *tclWashingMachine) Start() {
	fmt.Println("TCL WashingMachine Start...")
}

func (h *tclWashingMachine) Stop() {
	fmt.Println("TCL WashingMachine Stop...")
}

func (h *tclAirConditioner) Start() {
	fmt.Println("TCL AirConditioner Start...")
}

func (h *tclAirConditioner) Stop() {
	fmt.Println("TCL AirConditioner Stop...")
}
