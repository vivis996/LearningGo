package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var john person
	fmt.Println(john)
	fmt.Printf("%+v", john)
}
