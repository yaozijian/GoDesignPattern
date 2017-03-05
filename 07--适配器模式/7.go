package main

import "m"

func main(){

	r := m.NewRobot()
	r.Speak()
	r.Move()

	d := m.NewDog()
	d.Cry()
	d.Run()

	r = m.NewAdaptor(d)
	r.Speak()
	r.Move()
}

