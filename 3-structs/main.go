package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p := person{name: "med", age: 12}
	p.print()

	p.updateName1("khalid")
	p.print()

	pp := &p
	pp.updateName2("ahmed")

	p.print()

}

func (p person) updateName1(nn string) {
	p.name = nn
}

func (p *person) updateName2(nn string) {
	(*p).name = nn
}

func (p person) print() {
	fmt.Println("name :", p.name, " Age : ", p.age)
}