 package m

type(
	Meal interface{
		Food() string
		Drink() string
		setFood(string)
		setDrink(string)
	}

	meal struct{
		food string
		drink string
	}
)

func (m *meal) Food() string { return m.food }
func (m *meal) Drink() string { return m.drink }

func (m *meal) setFood(f string) { m.food = f }
func (m *meal) setDrink(d string) { m.drink = d }

