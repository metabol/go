package main

import (
"fmt"
)


func isEven(x int){
  if ans := x%2; ans == 0 {
      fmt.Println(x, " is even")
  }else{
     fmt.Println(x," is odd")
  }
}

func main(){

isEven(555)
isEven(600)
isEven(700887+968)
isEven(7609)

}
