package async_test

import (
  "testing"
  "fmt"
  "time"
)

func service() string {
  time.Sleep(time.Millisecond * 50)
  return "service done."
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func AsyncService() chan string  {
  retChan := make(chan string, 1)
  // retChan := make(chan string)
  go func ()  {
    ret := service()
    fmt.Println("returned result.")
    retChan <- ret
    fmt.Println("service exited.")
  }()

  fmt.Println("go func executed.")
  return retChan
}

func TestCSP(t *testing.T)  {
  retChan := AsyncService()
  otherTask()
  fmt.Println(<-retChan)
  fmt.Println("TestCSP Done.")
}
