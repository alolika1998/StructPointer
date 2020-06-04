package main

import "fmt"

type Employee struct {
Name string
Id int
City string
Position string
}
func main() {
	emp3 := &Employee {
		Name : "Lisa",
		Id : 1,
		City : "Kolkata",
		Position : "CEO",
	}
	fmt.Println(*emp3)
	fmt.Println("Name:", (*emp3).Name)
	fmt.Println("Id:", (*emp3).Id)
	emp3.Id = 11
	fmt.Println("New Id:", (*emp3).Id)
	fmt.Println(*emp3)
}
