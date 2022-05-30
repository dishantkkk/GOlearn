package main

import "fmt"

func main(){
	x:=[6] string {"dishant","imran","ajay","himansh"}
	fmt.Println(x)
	var y [] string=x[1:4]
	var z [] string=x[0:2]
	fmt.Println("slice 1 Y",y)	
	fmt.Println("slice 2 Z",z)	
	y[2]="Brillio"
	fmt.Println("After updating slice ",y)	
	fmt.Println("after value of X",x)
	fmt.Println(len(y))
		
	w := append(y,z...)
	fmt.Println(w)
	w[5]="imli"
	fmt.Println(w)
}