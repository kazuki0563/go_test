package main

import"fmt"

func main(){
  prime := [100]int{2}
  index := 1
  var flag int = 0
  for n:=3; n<101; n++{
    for i:=2; i <= n; i++{
      if i==n{
         flag = 1
      }
      if n%i==0{
        break
      }
    }
    if flag == 1{
      prime[index] = n
      index++
      flag = 0
    }
  }
  for i:=0; i<index; i++{
  fmt.Println(prime[i]) 
  }

}
