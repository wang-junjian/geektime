package series

import "fmt"

func init()  {
  fmt.Println("init1")
}

func init()  {
  fmt.Println("init2")
}

func Sum(x int, y int) int {
  return x+y
}

func print(s string)  {
  fmt.Println("> ", s)
}
