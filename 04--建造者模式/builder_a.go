package m

type meal_a struct{
	builderBase
}

func NewMealA() MealBuilder{
	m := &meal_a{}
	m._buildDrink = m.buildDrink
	m._buildFood = m.buildFood
	return m
}

func (m *meal_a) buildDrink(){
	m.reset()
	m.m.setDrink("一杯可乐")
}

func (m *meal_a) buildFood(){
	m.reset()
	m.m.setFood("一盒薯条")
}

