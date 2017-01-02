package main

import "m"
import "fmt"

func main(){

	w := m.NewWaiter()
	
	p := w.Construct(m.NewMealA())
	fmt.Printf("套餐A: 饮料: %v 食物: %v\n",p.Drink(),p.Food())
	
	p = w.Construct(m.NewMealB())
	fmt.Printf("套餐B: 饮料: %v 食物: %v\n",p.Drink(),p.Food())
}


