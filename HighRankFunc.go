
package main 
import"fmt"

func sumInt(n,m int)int{
  a :=0
  for;n<=m;n++{
    a+=n
  }
  return a
}

func sq(x int)int{
  return x*x
}

func cu(x int)int{
  return x*x*x
}

func sumOf(f func(int) int,n,m int)int{
  a := 0
  for; n<=m; n++{
    a += f(n)
  }
  return a
}

func main(){
  fmt.Println(sumOf(sq, 1, 100))
  fmt.Println(sumOf(cu, 1, 100))
}
