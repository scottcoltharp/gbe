package structexample

import "fmt"

type mammal interface {
	setAge(age int)
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 45
	return &p
}

func setAgeNow(m mammal, age int) {
	m.setAge(age)
}

func (person *person) setAge(age int) {
	person.age = age
}

func StructExample() {

	//fmt.Println(person{"Bob", 20})

	//fmt.Println(person{name: "Alice", age: 30})

	//fmt.Println(person{name: "Fred"})

	//fmt.Println(&person{name: "Ann", age: 40})

	//fmt.Println(newPerson("Jon"))

	ptrStu := newPerson("Scott")
	setAgeNow(ptrStu, 75)
	fmt.Println("pointer: ", ptrStu.age)

	//s := person{name: "Sean", age: 50}
	//fmt.Println(s.name)

	//sp := &s
	//fmt.Println(sp.age)

	//sp.age = 51
	//fmt.Println(sp.age)

	//dog := struct {
	//	name   string
	//	isGood bool
	//}{
	//	"Rex",
	//	true,
	//}
	//fmt.Println(dog)
}
