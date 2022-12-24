package main

import "fmt"

type Person struct{
	name string
	sex byte
	age int
}

func setPerson(p *Person){
	p.name = "setName"
	p.age = 10
}

func changePerson(p *Person){
	p.name = "qing"
	p.sex = 'M'
}

func initPerson()*Person{
	p := new(Person)
	p.name = "qing"
	p.sex = 'F'
	return p
}

func main(){
	p1 := Person{name: "test"}
	setPerson(&p1)
	fmt.Println("p1 ==> ", p1)

	p2 := new(Person)
	changePerson(p2)
	fmt.Println("p2 ==> ", p2)

	p3:=initPerson()
	fmt.Println("p3 ==> ", p3)
}