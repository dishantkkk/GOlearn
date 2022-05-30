package main

import (
	"FMT"
	)

func main() {
	fmt.Println("Hello, World!");
	
	var x int=100
	fmt.Println(x);
	
	x=200;
	fmt.Println("x = ",x);
	
	var y="Hlo";
	fmt.Println(y);
	
	z:=true
	fmt.Println(z);
	z=false
	fmt.Println(z);
	
	const a=10
	fmt.Println(a);
	
	const b int
	b=44 //can'y be done
	fmt.Println(b);
	
	var i,j=300,"Dishant";
	fmt.Println("i & j are",i,j)
	fmt.Println("i= j=",j)
	
	test();
}

func test(){
	fmt.Println("Hello")
}