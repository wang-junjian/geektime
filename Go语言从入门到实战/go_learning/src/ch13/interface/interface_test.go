package interface_test

import (
  "testing"
  "fmt"
)

type Code string

type Programmer interface {
  writeHello() Code
}

type GoProgrammer struct {

}

func (p *GoProgrammer) writeHello() Code  {
  return "fmt.Println(\"Hello\")"
}

type JavaProgrammer struct {

}

func (p *JavaProgrammer) writeHello() Code {
  return "System.out.println(\"Hello\")"
}

func writeHello(p Programmer)  {
  fmt.Printf("type: %T, Code: %s\n", p, p.writeHello())
}

func TestPolymorphism(t *testing.T)  {
  var p Programmer
  p = new(GoProgrammer)
  writeHello(p)

  p = new(JavaProgrammer)
  writeHello(p)
}
