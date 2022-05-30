package main

import "fmt"

func main(){
	defer fmt.Println("hello")
	
	fmt.Println("world")
	
}