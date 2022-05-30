package main

import "fmt"

func main(){
	fmt.Println("start");
	i:=1
	for {
	fmt.Println(i)
    i++
    if i >= 5 {
        break
    }
	}
	fmt.Println("done");
		
	
}