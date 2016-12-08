package m

type (
	Company interface {
		GetTV() TV
		GetWashingMachine() WashingMachine
		GetAirConditioner() AirConditioner
	}

	company struct {
		f Factory
	}
)

func NewCompany(f Factory) Company {
	return &company{f: f}
}

func (c *company) GetTV() TV {
	return c.f.CreateTV()
}

func (c *company) GetWashingMachine() WashingMachine {
	return c.f.CreateWashingMachine()
}

func (c *company) GetAirConditioner() AirConditioner {
	return c.f.CreateAirConditioner()
}
