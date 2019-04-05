package error_test

import (
  "testing"
  "fmt"
  "errors"
)

var LessThanTwoError error = errors.New("n must be greater than 2")
var GreaterThanHundredError error = errors.New("n must be less than 100")
func GetNums(n int) ([]int, error)  {
  if n < 2 {
    return nil, LessThanTwoError
  }
  if n > 100 {
    return nil, GreaterThanHundredError
  }
  return []int{n, n, n}, nil
}

func CallGenNums(n int)  {
  v, err := GetNums(n)

  if err == LessThanTwoError {
    fmt.Println("Need a larger number")
  } else if err == GreaterThanHundredError {
    fmt.Println("Need a smaller number")
  }

  fmt.Println(v, err)
}
func TestGetNums(t *testing.T)  {
  CallGenNums(0)
  CallGenNums(200)
  CallGenNums(10)
}
