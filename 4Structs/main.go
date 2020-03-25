// package main

// import "fmt"

// type person struct {
// 	firstName   string
// 	lastName    string
// 	age         int
// 	contactInfo contactInfo
// }

// type contactInfo struct {
// 	address string
// 	email   string
// 	zipCode int
// }

// func (p person) print() {
// 	fmt.Printf("%+v\n", p)
// }

// func (p *person) updateName(newFirstName string) {
// 	(*p).firstName = newFirstName
// }

// func main() {
// 	alex := person{firstName: "Alex", lastName: "Anderson", age: 20}
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Eric",
// 		age:       22,
// 		contactInfo: contactInfo{
// 			address: "lalalalala",
// 			zipCode: 44000,
// 			email:   "test@test.fr",
// 		},
// 	}
// 	jim.updateName("Eric")
// 	jim.print()
// 	alex.print()
// 	name := "Steeve"
// 	fmt.Print(*&name)
// }

package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
	fmt.Println(name)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
	(*namePointer) = "dog"
}
