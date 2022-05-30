package main

import "fmt"

func main(){
	fmt.Println("start");
	
	i:=1
	for i<10 {
		defer fmt.Println(i)
		i++
	}
	fmt.Println("done");
	
	
		
	
}