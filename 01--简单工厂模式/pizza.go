package m

import "fmt"

type (
	Pizza interface {
		Prepare()
		Bake()
		Cut()
		Box()
	}

	cheesePizza    struct{}
	pepperoniPizza struct{}
	veggiePizza    struct{}

	SimplePizzaFactory struct{}
)

func (*cheesePizza) Prepare() { fmt.Println("Prepare CheesePizza......") }
func (*cheesePizza) Bake()    { fmt.Println("Bake CheesePizza......") }
func (*cheesePizza) Cut()     { fmt.Println("Cut CheesePizza......") }
func (*cheesePizza) Box()     { fmt.Println("Box CheesePizza......") }

func (*pepperoniPizza) Prepare() { fmt.Println("Prepare PepperoniPizza......") }
func (*pepperoniPizza) Bake()    { fmt.Println("Bake PepperoniPizza......") }
func (*pepperoniPizza) Cut()     { fmt.Println("Cut PepperoniPizza......") }
func (*pepperoniPizza) Box()     { fmt.Println("Box PepperoniPizza......") }

func (*veggiePizza) Prepare() { fmt.Println("Prepare VeggiePizza......") }
func (*veggiePizza) Bake()    { fmt.Println("Bake VeggiePizza......") }
func (*veggiePizza) Cut()     { fmt.Println("Cut VeggiePizza......") }
func (*veggiePizza) Box()     { fmt.Println("Box VeggiePizza......") }

func (*SimplePizzaFactory) CreatePizza(typ string) Pizza {
	switch typ {
	case "cheese":
		return new(cheesePizza)
	case "pepperoni":
		return new(pepperoniPizza)
	case "veggie":
		return new(veggiePizza)
	default:
		return nil
	}
}
