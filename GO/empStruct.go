package main

import "fmt"

type Employee struct {
	name string
	age int
	address string
	phone string
}

func display_emp(e Employee) {
	fmt.Println(e.name)
}


func main() {
	var emp1= Employee{"ABC",23,"A-6 Delhi","1212121212"}
	//var emp2= Employee{"DEF",24,"A-7 Goa"}
	//var emp3= Employee{"GHI",25,"A-8 Kolkata"}
	
	display_emp(emp1)
	
	fmt.Println(emp1)
	//fmt.Println(emp2)
	//fmt.Println(emp3)
}