package main

import (
"fmt"
"runtime"

)


//Sitch dont need break statements
func main(){
fmt.Println("Your OS is :")

  switch os := runtime.GOOS; os {
  case "darwin" :
    fmt.Println("OS X")
  case "linux" :
    fmt.Println("Linux")
  default :
     fmt.Println(os)
  }

}
