package client_test

import (
  "testing"
  "fmt"
  "ch15/series"
)

func TestPackage(t *testing.T)  {
  fmt.Println(series.Sum(1, 2))

  // 小写字母开头的函数不能在包外面调用
  // series.print("hello")
}
