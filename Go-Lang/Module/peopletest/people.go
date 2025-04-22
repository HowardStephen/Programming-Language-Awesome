package people

import "fmt"

type People struct {
	name string
	age  int
}

func (p *People) DefinePeople(name string, age int) {
	p.name = name
	p.age = age
}

func (p *People) ShowPeopleInfo() {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}
