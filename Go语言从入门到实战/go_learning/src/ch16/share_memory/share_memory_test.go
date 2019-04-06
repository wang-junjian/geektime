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

func TestCounterWaitGroup(t *testing.T)  {
  var mutex sync.Mutex
  var waitGroup sync.WaitGroup

  counter := 0
  for i := 0; i < 5000; i++ {
    waitGroup.Add(1)
    go func() {
      defer func ()  {
        mutex.Unlock()
        waitGroup.Done()
      }()

      mutex.Lock()
      counter++
    }()
  }
  waitGroup.Wait()
  
  fmt.Println("counter = ", counter)
}
