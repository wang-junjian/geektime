package singleton_test

import (
  "fmt"
  "sync"
  "testing"
  "unsafe"
)

type Singleton struct {

}

var singleton *Singleton
var once sync.Once

func GetSingletonInstance() *Singleton {
  once.Do(func ()  {
    fmt.Println("GetSingletonInstance() once run.")
    singleton = new(Singleton)
  })

  return singleton
}

func TestGetSingletonInstance(t *testing.T)  {
  var wg sync.WaitGroup

  for i:=0; i<10; i++ {
    wg.Add(1)
    go func(i int) {
      s1 := GetSingletonInstance()
      fmt.Println(i, unsafe.Pointer(s1))

      wg.Done()
    }(i)
  }

  wg.Wait()
}
