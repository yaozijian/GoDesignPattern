package m

type (
	newYorkCheesePizza struct{ pizzaBase }
	newYorkPizzaStore  struct{ pizzaStore }
)

func NewNewYorkPizzaStore() PizzaStore {
	store := new(newYorkPizzaStore)
	store._createPizza = store.createPizza
	return store
}

func (s *newYorkPizzaStore) createPizza(typ string) (p Pizza) {
	switch typ {
	case "cheese":
		p = newNewYorkCheesePizza()
	}
	return
}

func newNewYorkCheesePizza() Pizza {
	return &newYorkCheesePizza{
		pizzaBase: pizzaBase{
			name:     "NewYork Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sause:    "Marinara Sauce",
			toppings: []string{"Crated Reggiano Cheese"},
		},
	}
}
