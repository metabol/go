package main

import (
"fmt"
)


type Shape struct{
  L int
  W int
}

func main(){

fmt.Println(Shape{0,0})

r := &Shape
fmt.Println(r)

//(*r).L = 5    //Access by pointer
//r.W = 6       //Direct access permitted 

//fmt.Println(*r)


}
