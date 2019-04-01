package operator_test

import (
  "testing"
  "fmt"
)

func TestOperator(t *testing.T)  {
  var a = [...]int{1, 2, 3, 4}
  var b = [...]int{1, 2, 3, 4}
  var c = [...]int{1, 2, 5, 4}

  fmt.Println(a==b)
  fmt.Println(a==c)
}
