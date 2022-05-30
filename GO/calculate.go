package main

import "fmt"

func calculate(num1 int, num2 int) (int, int) {
sum := num1+num2
diff := num1-num2

return sum,diff

}

func main(){
	x,y :=30,20
	sum, diff := calculate(x,y)
	
	fmt.Println(sum)
	fmt.Println(diff)
	
	defer calculate1(x,y)
	fmt.Println("hii")
}

func calculate1(num1 int, num2 int) {
sum := num1+num2
diff := num1-num2
fmt.Println(sum)
fmt.Println(diff)
}
