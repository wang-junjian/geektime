package panic_test

import (
  "testing"
  "fmt"
  "os"
)

func TestPanic(t *testing.T)  {
  defer func ()  {
    fmt.Println("TestPanic Clear")
  }()

  fmt.Println("Start")

  panic("Error")
}

func TestExit(t *testing.T)  {
  defer func ()  {
    fmt.Println("TestExit Clear")
  }()

  fmt.Println("Start")

  os.Exit(-1)
}
