package m

type(
	MealBuilder interface{
		getMeal() Meal
		buildDrink()
		buildFood()
	}

	builderBase struct{
		m *meal
		_buildDrink func()
		_buildFood func()
	}
)

func (b *builderBase) reset(){
	if b.m == nil{
		b.m = new(meal)
	}
}

func (b *builderBase) buildDrink(){
	b._buildDrink()
}

func (b *builderBase) buildFood(){
	b._buildFood()
}

func (b *builderBase) getMeal() Meal{
	m := b.m
	b.m = nil
	return m
}


