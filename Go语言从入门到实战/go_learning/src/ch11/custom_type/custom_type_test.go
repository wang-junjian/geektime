package custom_type

import (
  "testing"
  "fmt"
  "time"
)

func square(n int) int  {
  time.Sleep(time.Second)
  return n*n
}

type square_func_type func(n int) int

func TimeMeasure(f square_func_type) square_func_type {
  return func (n int) int  {
    begin := time.Now()
    ret := f(n)
    fmt.Println(time.Since(begin).Seconds())
    return ret
  }
}

func TestCustomType(t *testing.T)  {
  timeMeasureSquare := TimeMeasure(square)
  fmt.Println(timeMeasureSquare(10))
}
