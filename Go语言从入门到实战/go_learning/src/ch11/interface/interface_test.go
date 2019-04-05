package interface_test

import (
  "testing"
  "fmt"
)

type Programmer interface {
  Write() string
}

type GoProgrammer struct {

}

func (p *GoProgrammer) Write() string {
  return "Hello Go Programmer"
}

func TestInterface(t *testing.T)  {
  var p Programmer
  p = new(GoProgrammer)
  fmt.Println(p.Write())
}
