package main

import (
"fmt"
)



func main(){
k := 5
p := &k
fmt.Println(p)
fmt.Println(*p)
*p = 45
fmt.Println(p)
fmt.Println(*p)


}
