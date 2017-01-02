package main

import "fmt"
import "m"

func main(){
	
	p1 := m.NewSingleton()
	fmt.Println("Address p1:",p1.Address())
	
	p2 := m.NewSingleton()
	fmt.Println("Address p2:",p2.Address())
	
	p3 := m.NewSingleton()
	fmt.Println("Address p3:",p3.Address())
}

