package m

type (
	chicagoCheesePizza struct{ pizzaBase }
	chicagoPizzaStore  struct{ pizzaStore }
)

func NewChicagoPizzaStore() PizzaStore {
	store := new(chicagoPizzaStore)
	store._createPizza = store.createPizza
	return store
}

func (s *chicagoPizzaStore) createPizza(typ string) (p Pizza) {
	switch typ {
	case "cheese":
		p = newChicagoCheesePizza()
	}
	return
}

func newChicagoCheesePizza() Pizza {
	return &chicagoCheesePizza{
		pizzaBase: pizzaBase{
			name:     "Chicago Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sause:    "Marinara Sauce",
			toppings: []string{"Crated Reggiano Cheese"},
		},
	}
}
