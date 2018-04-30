package main

import (
"fmt"
)


//Function that returns two string values
func swap(a,b string)(string,string){
 return b,a
}



func main(){

a, b := swap("one","two")
fmt.Println(a," ",b)
}
