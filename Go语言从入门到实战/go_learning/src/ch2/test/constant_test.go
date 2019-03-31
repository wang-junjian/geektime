package try_test

import (
  "testing"
  "fmt"
)

func TestConstant(t *testing.T)  {
  const (
    One = 1 + iota
    Two
    Three
  )

  fmt.Println(One, Two, Three)

  const (
    Read = 1 << iota
    Write
    Execute
  )

  flag:=7 //0111
  fmt.Println(flag&Read==Read, flag&Write==Write, flag&Execute==Execute)
}
