package m

import "fmt"

type(
	adaptor struct{
		dog Dog
	}
)

func NewAdaptor(d Dog) Robot{
	return &adaptor{dog:d}
}

func (a *adaptor) Speak(){
	fmt.Println("适配器用狗模拟机器人Speak")
	a.dog.Cry()
}

func (a *adaptor) Move(){
	fmt.Println("适配器用狗模拟机器人Move")
	a.dog.Run()
}


