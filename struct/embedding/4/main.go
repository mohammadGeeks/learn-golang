package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   *Address
	Contact   Contact
}

type Address struct {
	Street string
	City   string
	Zip    string
}

type Contact struct {
	Email string
	Phone string
}

func main() {
	homeAddress := &Address{
		Street: "123 Main St",
		City:   "Tehran",
		Zip:    "123456",
	}

	personalContact := Contact{
		Email: "john.doe@example.com",
		Phone: "123-456-7890",
	}

	john := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address:   homeAddress,
		Contact:   personalContact,
	}

	fmt.Println(john.FirstName, john.LastName)
	fmt.Println(john.Contact)
	fmt.Println(john.Address)

	mohammad := john
	fmt.Println(mohammad.Address)

	mohammad.Address.City = "rasht"
	fmt.Println(mohammad.Address)
	fmt.Println(john.Address)
}
