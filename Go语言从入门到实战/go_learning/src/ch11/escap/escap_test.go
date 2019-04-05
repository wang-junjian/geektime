package escap_test

import (
  "testing"
  "fmt"
  "unsafe"
)

type Employee struct {
  name string
  age int
}

func (e *Employee) String() string {
  fmt.Printf("  Employee Address %x, name Address %x\n", unsafe.Pointer(e), unsafe.Pointer(&e.name))
  return fmt.Sprintf("(%T) name:%s, age:%d", e, e.name, e.age)
}

// func (e Employee) String() string {
//   fmt.Printf("  Employee Address %x, name Address %x\n", unsafe.Pointer(&e), unsafe.Pointer(&e.name))
//   return fmt.Sprintf("(%T) name:%s, age:%d", e, e.name, e.age)
// }

func TestStructDefine(t *testing.T)  {
  e1 := Employee{"wjj", 38}
  fmt.Println(e1)

  e2 := Employee{age:10, name:"wrj"}
  fmt.Println(e2)

  e3 := new(Employee)
  e3.name = "xxx"
  e3.age = 18
  fmt.Println(e3)

  fmt.Println(fmt.Sprintf("e1:%T, e2:%T, e3:%T", e1, e2, e3))

  fmt.Printf("1 Employee Address %x, name Address %x\n", unsafe.Pointer(&e1), unsafe.Pointer(&e1.name))
  fmt.Println(e1.String())
  fmt.Printf("2 Employee Address %x, name Address %x\n", unsafe.Pointer(&e2), unsafe.Pointer(&e2.name))
  fmt.Println(e2.String())
  fmt.Printf("3 Employee Address %x, name Address %x\n", unsafe.Pointer(e3), unsafe.Pointer(&e3.name))
  fmt.Println(e3.String())
}
