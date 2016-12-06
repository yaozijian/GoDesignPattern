package main

import "m"
import "fmt"

func main() {

	fmt.Println("----- Joel need chicago deep dish pizza -----")
	store1 := m.NewChicagoPizzaStore()
	p1 := store1.OrderPizza("cheese")
	fmt.Println("Joel ordered a ", p1.Name())

	fmt.Println("\n\n----- Ethan need newyork style pizza -----")
	store2 := m.NewNewYorkPizzaStore()
	p2 := store2.OrderPizza("cheese")
	fmt.Println("Ethan ordered a ", p2.Name())
}
