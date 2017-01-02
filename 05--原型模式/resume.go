package m

import "fmt"

type(
	Resume struct{
		name string
		birth string
		sex string
		school string
		timeRange string
		company string
	}
)

func NewResume(name string) *Resume{
	return &Resume{name:name}
}

func (r *Resume) SetPersonInfo(birth,sex,school string){
	r.birth = birth
	r.sex = sex
	r.school = school
}

func (r *Resume) SetWorkExperience(timerange,company string){
	r.timeRange = timerange
	r.company = company
}

func (r *Resume) String() string{
	return fmt.Sprintf(
		"姓名: %v 生日: %v 性别: %v\n学校: %v\n公司: %v 时间: %v\n",
		r.name,r.birth,r.sex,r.school,r.company,r.timeRange,
	)
}

func (r *Resume) Clone() *Resume{
	return r.clone1()
}

func (r *Resume) clone1() *Resume{
	return &Resume{
		name: r.name,
		sex: r.sex,
		birth: r.birth,
		school: r.school,
		company: r.company,
		timeRange: r.timeRange,
	}
}

