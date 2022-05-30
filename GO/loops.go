package main

import (
	"FMT"
	)

func main() {
	fmt.Println("Hello again!");
		
	var a int
	for a=1;a<=5;a++ {
		fmt.Println(a)
	}
	var x int=10
	if x%2==0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}
	var r,p=1,2
	fmt.Println(r+p)	
	i,j :=2,1
	switch i+j {
		case 1 : fmt.Println("i=",i)
		case 3 : fmt.Println("j=",j)
		default : fmt.Println("default")
		}
		
	var arr [5] int
	for i=0;i<5;i++ {
		arr[i]=i
	}
	fmt.Println(arr)
	fmt.Println(len(arr))
	
	newArray := [] int {10,20,30,40,50}
	fmt.Println(newArray)
	fmt.Println(newArray[5])
	
}