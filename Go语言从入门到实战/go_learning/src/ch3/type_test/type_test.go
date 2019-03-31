package type_test

import (
  "testing"
  "fmt"
  "math"
)

func TestPreDefined(t *testing.T)  {
  fmt.Println(math.MinInt64, math.MaxInt64)
  fmt.Println(math.MaxFloat64)
  fmt.Println(math.MaxUint32)
}

func TestPointer(t *testing.T) {
  a := 1
  aptr := &a
  fmt.Println(a, aptr)

  // aptr = aptr + 1 //invalid operation: aptr + 1 (mismatched types *int and int)
  // aptr = aptr + aptr //invalid operation: aptr + aptr (operator + not defined on pointer)
}

func TestImplictCast(t *testing.T) {
  var a int = 1
  var b float64 = 2.2
  fmt.Println(a, b)

  // b = a //cannot use a (type int) as type float64 in assignment
  b = float64(a)

  type MyInt int
  var c MyInt
  fmt.Println(c)
  // c = a //cannot use a (type int) as type MyInt in assignment
}
