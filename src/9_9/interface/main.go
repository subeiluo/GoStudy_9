package main

import "fmt"

type Employee interface {
	Calc() float32
}
type EmployeeMgr struct {
	employeelist []Employee
}

type Developer struct {
	Name string
	Base float32
}
func (d *Developer)Calc()float32{
	return d.Base
}

type PM struct {
	Name string
	Base float32
	Option float32
}

func (p *PM)Calc()float32  {
	return p.Base+p.Option
}

type YY struct {
	Name string
	Base float32
	Option float32
	Rate float32
}

func (p *YY)Calc()float32  {
	return p.Base+p.Option*p.Rate
}

func(e *EmployeeMgr)Calc()float32{
	var sum float32
	for _,v:=range e.employeelist{
		sum+=v.Calc()
	}
	return sum
}

func (e *EmployeeMgr)AddEmpoyee(d Employee)  {
	e.employeelist =append(e.employeelist,d)
}
func main()  {
	var e  =&EmployeeMgr{}
	dev :=&Developer{
		Name:"develop",
		Base:10000,
	}
	e.AddEmpoyee(dev)

	pm :=&PM{
		Name:"pm",
		Base:8000,
		Option:2389.5,
	}
	e.AddEmpoyee(pm)

	yy :=&YY{
		Name:"yy",
		Base:9980,
		Option:3000,
		Rate:0.7,
	}
	e.AddEmpoyee(yy)

	sum :=e.Calc()
	fmt.Printf("sum:%f\n",sum)
}