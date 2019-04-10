package cancel_task_test

import (
  "testing"
  "fmt"
  "time"
)

func isCancelled(cancelChan chan struct{}) bool {
  select {
  case <-cancelChan:
    return true
  default:
    return false
  }
}
func TestCancle(t *testing.T)  {
  cancelChan := make(chan struct{}, 0)
  for i:=0; i<5; i++ {
    go func(i int, cancelChan chan struct{}) {
      for {
        if isCancelled(cancelChan) {
          break
        }

        time.Sleep(time.Millisecond * 5)
      }

      fmt.Println(i, "Cancelled.")
    }(i, cancelChan)
  }

  close(cancelChan)
  time.Sleep(time.Second * 1)
}
