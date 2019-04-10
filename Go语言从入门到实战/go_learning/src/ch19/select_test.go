package select_test

import (
  "testing"
  "fmt"
  "time"
)

func serviceWith100MS() int {
  time.Sleep(time.Millisecond * 100)
  return 100
}

func serviceWith500MS() int {
  time.Sleep(time.Millisecond * 500)
  return 500
}

func AsyncService(serviceFunc func() int) chan int {
  retChan := make(chan int, 1)
  go func() {
    ret := serviceFunc()
    retChan <- ret
  }()

  return retChan
}

func TestSelect(t *testing.T)  {
  retChan1 := AsyncService(serviceWith100MS)
  retChan2 := AsyncService(serviceWith500MS)

  select {
  case ret := <-retChan1:
    fmt.Println("select Async call serviceWith100MS returned ", ret)
  case ret := <-retChan2:
    fmt.Println("select Async call serviceWith500MS returned ", ret)
  case <-time.After(time.Second * 1):
    fmt.Println("time out.")
  // default:
  //   fmt.Println("select default!")
  }

  fmt.Println("Done.")
}
