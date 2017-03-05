package m

import "fmt"

type(
	Robot interface{
		Speak()
		Move()
	}

	bioRobot struct{}
)

func NewRobot() Robot{
	return &bioRobot{}
}

func (b *bioRobot) Speak(){
	fmt.Println("仿生机器人发声...")
}

func (b *bioRobot) Move(){
	fmt.Println("仿生机器人移动...")
}

