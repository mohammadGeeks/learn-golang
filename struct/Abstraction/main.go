package main

import "fmt"

const (
	BaseSalary       = 5600000
	ExtraHourRate    = 90000
	HourlySalaryRate = 110000
	ShiftSalaryRate  = 80000
	TaxRate          = 0.09
)

type Employee interface {
	SalaryCalculator(hours float64) (salary float64, tax float64)
}

type FullTimeEmployee struct {
	FullName   string
	ExtraHours float64
}

type PartTimeEmployee struct {
	FullName      string
	PartTimeHours float64
}

func (employee FullTimeEmployee) SalaryCalculator(extraHours float64) (salary float64, tax float64) {
	salary = BaseSalary + (ExtraHourRate*extraHours)*1.4
	tax = salary * TaxRate
	salary += tax
	return
}

func (employee PartTimeEmployee) SalaryCalculator(hours float64) (salary float64, tax float64) {
	salary = HourlySalaryRate * hours
	tax = salary * TaxRate
	salary += tax
	return
}

func main() {

	fullTimeEmployees := []FullTimeEmployee{
		{FullName: "Pejman Rezaee", ExtraHours: 40},
		{FullName: "Maryam Hosseini", ExtraHours: 120},
	}

	partTimeEmployees := []PartTimeEmployee{
		{FullName: "Milad Hassani", PartTimeHours: 100},
		{FullName: "Maryam Rezaee", PartTimeHours: 87},
	}

	var employee Employee = fullTimeEmployees[0]

	salary, tax := employee.SalaryCalculator(float64(fullTimeEmployees[0].ExtraHours))
	fmt.Printf("Employee: %v\tSalary: %f\tTax: %f\n", employee, salary, tax)

	employee = fullTimeEmployees[1]
	salary, tax = employee.SalaryCalculator(float64(fullTimeEmployees[1].ExtraHours))
	fmt.Printf("Employee: %v\tSalary: %f\tTax: %f\n", employee, salary, tax)

	employee = partTimeEmployees[0]
	salary, tax = employee.SalaryCalculator(float64(partTimeEmployees[0].PartTimeHours))
	fmt.Printf("Employee: %v\tSalary: %f\tTax: %f\n", employee, salary, tax)

	employee = partTimeEmployees[1]
	salary, tax = employee.SalaryCalculator(float64(partTimeEmployees[1].PartTimeHours))
	fmt.Printf("Employee: %v\tSalary: %f\tTax: %f\n", employee, salary, tax)

}
