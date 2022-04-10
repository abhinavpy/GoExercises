package main

import "fmt"

/*
	Go's structs are typed collections of fields. 
	They're useful for grouping data together to form records.

	This person struct type has name and age fields.

	newPerson constructs a new person struct with the given name.
	You can safely return a pointer to local variable as a local variable will
	survive the scope of the function.

	This syntax creates a new struct.

	You can name the fields when initializing a struct.
	Omitted fields will be zero valued.

	An & prefix yields a pointer to the struct.

	It's idiomatic to encapsulate new struct creation in constructor functions.
	Access struct fields with a dot.

	You can also use dots with struct pointers - the pointers are automatically dereferenced.
	Structs are mutable.
*/

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
