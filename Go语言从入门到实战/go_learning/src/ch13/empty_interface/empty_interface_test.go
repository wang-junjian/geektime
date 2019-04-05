package empty_interface

import (
  "testing"
  "fmt"
)

func DoSomething(p interface{})  {
  if v, ok := p.(int); ok {
    fmt.Println("Integer", v)
    return
  }

  if v, ok := p.(string); ok {
    fmt.Println("String", v)
    return
  }

  fmt.Println("Unknow Type")
}

func DoSomething2(p interface{})  {
  switch v := p.(type) {
  case int:
    fmt.Println("Integer", v)
  case string:
    fmt.Println("String", v)
  default:
    fmt.Println("Unknow Type")
  }
}

func TestEmptyInterface(t *testing.T)  {
  //方法1
  DoSomething(10)
  DoSomething("hello")
  DoSomething(11.23)

  //方法2
  DoSomething2(10)
  DoSomething2("hello")
  DoSomething2(11.23)
}
