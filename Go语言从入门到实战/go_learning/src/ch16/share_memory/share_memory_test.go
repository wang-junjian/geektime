package share_memory_test

import (
  "testing"
  "fmt"
  "sync"
  "time"
)

func TestCounter(t *testing.T)  {
  counter := 0
  for i := 0; i < 5000; i++ {
    go func() {
      counter++
    }()
  }

  time.Sleep(time.Microsecond * 50)
  fmt.Println("counter = ", counter)
}

func TestCounterThreadSafe(t *testing.T)  {
  var mutex sync.Mutex

  counter := 0
  for i := 0; i < 5000; i++ {
    go func() {
      // mutex.Lock()
      // counter++
      // mutex.Unlock()

      defer func ()  {
        mutex.Unlock()
      }()

      mutex.Lock()
      counter++
    }()
  }

  time.Sleep(time.Microsecond * 50)
  fmt.Println("counter = ", counter)
}
