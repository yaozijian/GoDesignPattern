package m

type(
	Waiter interface{
		Construct(MealBuilder) Meal
	}
	
	waiter struct{
	}
)

func NewWaiter() Waiter{
	return &waiter{}
}

func (w *waiter) Construct(b MealBuilder) Meal{
	b.buildDrink()
	b.buildFood()
	return b.getMeal()	
}

