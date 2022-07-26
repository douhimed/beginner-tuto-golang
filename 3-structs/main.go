package main

import "fmt"

type contactInfos struct {
	tel string
	email string
}

type person struct {
	name string
	age  int
	contactInfos
}

func main() {

	p := person{name: "med", age: 12, contactInfos: contactInfos{tel: "0653251578", email:"test@test.com"}}
	p.print()
	p.updateName("ahmed")
	p.print()
	pPoint := &p
	pPoint.updateAge(25)
	p.print()
}

func (p person) updateName(n string) {
	p.name = n
}

func (p *person) updateAge(a int) {
	(*p).age = a
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}