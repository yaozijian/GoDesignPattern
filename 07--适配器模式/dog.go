package m

import "fmt"

type(
	Dog interface{
		Cry()
		Run()
	}

	habaDog struct{}
)

func NewDog() Dog{
	return &habaDog{}
}

func (h *habaDog) Cry(){
	fmt.Println("哈巴狗叫...")
}

func (h *habaDog) Run(){
	fmt.Println("哈巴狗跑...")
}

