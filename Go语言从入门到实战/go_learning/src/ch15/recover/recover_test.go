package panic_test

import (
  "testing"
  "fmt"
)

func TestRecover(t *testing.T)  {
  defer func ()  {
    if err := recover(); err!=nil {
      fmt.Println(err)
    }
  }()

  panic("Error")
}
