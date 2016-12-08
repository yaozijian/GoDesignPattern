package m

type (
	TV interface {
		Open()
		Close()
	}

	WashingMachine interface {
		Start()
		Stop()
	}

	AirConditioner interface {
		Start()
		Stop()
	}

	Factory interface {
		CreateTV() TV
		CreateWashingMachine() WashingMachine
		CreateAirConditioner() AirConditioner
	}
)
