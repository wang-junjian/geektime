package array_test

import (
  "testing"
  "fmt"
)

func TestArrayInit(t *testing.T)  {
  var a [3]int
  fmt.Println(a)

  b := [3]int{1, 2, 3}
  fmt.Println(b)

  c := [...]int{1, 2, 3, 4, 5}
  fmt.Println(c)

  d := [2][3]int{{1,2,3}, {4,5,6}}
  fmt.Println(d)
}

func TestArrayTravel(t *testing.T)  {
  a := [...]int{1, 2, 3, 4}

  for i:=0; i<len(a); i++ {
    fmt.Println(a[i])
  }

  for i, n := range a {
    fmt.Println(i, n)
  }

  for _, n := range a {
    fmt.Println(n)
  }
}
