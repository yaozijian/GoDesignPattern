package m

import "fmt"
import "sync"

type(
	Singleton interface{
		Address() string	
	}

	singleton struct{
	}
)

var(
	initonce sync.Once
	ptr *singleton
)

func NewSingleton() Singleton{
	initonce.Do(func(){ ptr = new(singleton) })
	return ptr
}

func (x *singleton) Address() string{
	return fmt.Sprintf("%p",x)
}

