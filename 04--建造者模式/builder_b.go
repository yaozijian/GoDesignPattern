package m

type meal_b struct{
	builderBase
}

func NewMealB() MealBuilder{
	m := &meal_b{}
	m._buildDrink = m.buildDrink
	m._buildFood = m.buildFood
	return m
}

func (m *meal_b) buildDrink(){
	m.reset()
	m.m.setDrink("一杯柠檬汁")
}

func (m *meal_b) buildFood(){
	m.reset()
	m.m.setFood("三个鸡翅")
}

