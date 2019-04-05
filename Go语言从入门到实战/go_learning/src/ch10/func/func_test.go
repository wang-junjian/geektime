package func_test

import (
  "testing"
  "fmt"
  "time"
)

func funcMultiReturnValues() (int, int) {
  return 10, 20
}

func TestFuncMultiReturnValues(t *testing.T)  {
  x, y := funcMultiReturnValues()
  fmt.Println(x, y)
}

func sleep()  {
  time.Sleep(time.Second)
}

func TestTimeMeasure(t *testing.T) {

  time_measure := func(f func())  {
    start := time.Now()
    f()
    tm := time.Since(start).Seconds()
    fmt.Println("user time: ", tm)
  }

  time_measure(sleep)
  fmt.Println()
}

func Clear() {
  fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T)  {
  defer Clear()
  fmt.Println("Start")
}

func TestDefer1(t *testing.T) {
  defer func (ops ...int)  {
    count := 0
    for _, op := range ops {
      count += op
    }

    fmt.Println("Count: ", count, time.Now())
  }(1, 2, 3, 4)

  fmt.Println("Begin")
  panic("error")
}
