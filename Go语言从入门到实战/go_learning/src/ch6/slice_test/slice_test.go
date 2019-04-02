package slice_test

import (
  "testing"
  "fmt"
  "reflect"
)

func TestSliceInit(t *testing.T)  {
  var a []int
  fmt.Println(a, len(a), cap(a))

  b := []int{1,2,3}
  fmt.Println(b, len(b), cap(b))

  b = append(b, 4)
  fmt.Println(b, len(b), cap(b))

  var c []int
  for i:=1; i<=10; i++ {
    c = append(c, i)
    fmt.Println(len(c), cap(c), reflect.TypeOf(c), c)
  }
}

func TestSliceShare(t *testing.T) {
  n := [...]int{1,2,3,4,5,6,7,8,9,10}
  fmt.Println(n)

  n1 := n[1:5]
  n2 := n[3:7]
  fmt.Println(n1, n2)

  n2[0] = 777
  fmt.Println(n, n1, n2)
}
