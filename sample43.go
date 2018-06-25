
package main

import "fmt"

func pow(x,y int)int{
  if y == 1{
    return x
  }else{
    return x*pow(x,y-1)
  }
}

func main(){
  for i:=0; i<15; i++{
    fmt.Println(i,"^2:",pow(i,2))
  }
}
