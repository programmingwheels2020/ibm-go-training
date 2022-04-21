package main

import "fmt"

type Student struct {
	name    string
	age     int
	subject string
}

type Employee struct {
	student Student
	empId   int
}

func newStudent(name string, age int, subject string) Student {
	return Student{name: name, age: age, subject: subject}
}

func main() {
	s := Student{name: "Lenin", subject: "Computer Science"}
	fmt.Println(s)
	s1 := Student{"Lenin", 67, "Computer science"}
	fmt.Println(s1)
	s2 := newStudent("Lenin", 67, "Comp Science")
	fmt.Println(s2.name)

	s4 := &s
	s4.age = 56
	fmt.Println(s)
	e1 := Employee{student: s, empId: 3223}
	fmt.Println(e1.student.name)

}
