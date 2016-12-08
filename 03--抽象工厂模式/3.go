package main

import "m"
import "fmt"

func main() {

	fmt.Println("----- Haier factory -----")

	c := m.NewCompany(m.NewHaierFactory())
	c.GetTV().Open()
	c.GetAirConditioner().Start()

	fmt.Println("\n\n----- TCL factory -----")

	c = m.NewCompany(m.NewTCLFactory())
	c.GetTV().Open()
	c.GetAirConditioner().Start()
}
