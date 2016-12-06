package m

import "fmt"

type (
	PizzaStore interface {
		OrderPizza(typ string) Pizza
	}

	Pizza interface {
		Name() string
		prepare()
		bake()
		cut()
		box()
	}

	//-----

	pizzaStore struct {
		_createPizza func(string) Pizza
	}

	pizzaBase struct {
		name     string
		dough    string
		sause    string
		toppings []string
	}
)

func (s *pizzaStore) OrderPizza(typ string) Pizza {
	p := s.createPizza(typ)
	p.prepare()
	p.bake()
	p.cut()
	p.box()
	return p
}

func (s *pizzaStore) createPizza(typ string) Pizza {
	return s._createPizza(typ)
}

//-----

func (p *pizzaBase) prepare() {
	fmt.Println("Preparing ", p.name)
	fmt.Println("Tossing dough")
	fmt.Println("Adding sause")
	fmt.Println("Adding toppings")
	for _, t := range p.toppings {
		fmt.Println("      ", t)
	}
}

func (p *pizzaBase) Name() string {
	return p.name
}

func (p *pizzaBase) bake() {
	fmt.Println("Bake for 25 minutes at 350℃")
}

func (p *pizzaBase) cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *pizzaBase) box() {
	fmt.Println("Place pizza in offical PizzaStore box")
}
