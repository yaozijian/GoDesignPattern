
package main

import "m"
import "fmt"

func main(){

	p := m.NewResume("孙悟空")
	p.SetPersonInfo("221 B.C","男","花果山")
	p.SetWorkExperience("210 B.C","无量方寸山")
	
	c := p.Clone()

	fmt.Println("----- A -----")
	fmt.Println(p.String())
	
	fmt.Println("----- B -----")
	fmt.Println(c.String())
}

