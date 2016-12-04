package main

import "m"

type pizzaStore struct {
	factory *m.SimplePizzaFactory
}

func newPizzaStore() *pizzaStore {
	return &pizzaStore{factory: new(m.SimplePizzaFactory)}
}

func (store *pizzaStore) orderPizza(typ string) (pizza m.Pizza) {

	pizza = store.factory.CreatePizza(typ)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return
}

func main() {
	store := newPizzaStore()
	store.orderPizza("cheese")
}
