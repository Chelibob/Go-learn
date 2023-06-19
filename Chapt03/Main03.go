package main

import "fmt"

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (e employee) getInfo() {
	fmt.Printf("Сотрудник %s\nВозраст: %d\nЗарплата: %d\n", e.name, e.age, e.salary)
}

func setName(e *employee, name string) {
	e.name = name
}

func (e *employee) setName(name string) {
	e.name = name
}

func main() {
	employee1 := newEmployee("Antoxa", "M", 25, 90000)
	employee2 := newEmployee("Alex", "W", 35, 45000)

	employee1.getInfo()
	employee2.getInfo()

	setName(&employee1, "Tolik")
	employee1.getInfo()
	employee2.setName("Boba")
	employee2.getInfo()
}
