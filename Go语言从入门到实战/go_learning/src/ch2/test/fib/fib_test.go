package fib

import (
  "testing"
  "fmt"
)


func TestFib(t *testing.T) {
  a:=1
  b:=1

  fmt.Print(a, " ")
  fmt.Print(b)
  for i:=3; i<=10; i++ {
    tmp:=a
    a = b
    b = tmp+a

    fmt.Print(" ", b)
  }

  fmt.Println()
}

func TestExchange(t *testing.T) {
  a:=1
  b:=2

  fmt.Println("a, b = ", a, b)
  a, b = b, a
  fmt.Println("a and b exchange: ", a, b)
}
