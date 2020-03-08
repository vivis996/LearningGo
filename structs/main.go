package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jane := person{
		firstName: "Jane",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email:   "jane.doe@example.com",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", jane)
}
