package main

import(
"fmt"
)

//NOte that the type is specified once bcos both params are expected to be int
func sub(x,y int) int{
  return x-y
}

func main(){
  fmt.Println( sub(999,888))
}
