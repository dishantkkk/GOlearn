package main

import "fmt"

func main() {
   students := map[string]int{
    "Imran":24,
    "Ajay":23,
    "Himansh":25,
   }

   students["Sumit"]=12
   for name := range students {
      fmt.Println("Age of",name,"is",students[name])
   }
   age, ok := students["Sumit"]

   if(ok){
      if(age>15){
         fmt.Println("Greater", age)
      } else {
         fmt.Println("Less",age)
      }
   } else {
      fmt.Println("Less")
   }
}
