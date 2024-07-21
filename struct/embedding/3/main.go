package main

import "fmt"

type ContactInfo interface {
	GetDetails() string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
	Contact   ContactInfo
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

func (c Contact) GetDetails() string {
	return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
}

func main() {

	homeAddress := Address{
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

	fmt.Printf("Name: %s %s\n", john.FirstName, john.LastName)
	fmt.Printf("Age: %d\n", john.Age)
	fmt.Printf("Address: %s, %s, %s\n", john.Address.Street, john.Address.City, john.Address.Zip)
	fmt.Printf("Contact: %s\n", john.Contact.GetDetails())
}
