package string_test

import (
  "testing"
  "fmt"
  "strings"
  "strconv"
)

func TestInit(t *testing.T)  {
  s := "hello"
  fmt.Println(s, len(s))

  for i, c := range s {
    fmt.Println(i, c)
  }

  s = "中华人民共和国"
  fmt.Println(s, len(s))

  s1 := []rune(s)
  fmt.Println(s1, len(s1))

  for _, c := range s {
    fmt.Println(c)
  }
}

func TestStringFunc(t *testing.T) {
  s := "A,B,C"
  ss := strings.Split(s, ",")
  fmt.Println(s, ss)

  s = strings.Join(ss, "-")
  fmt.Println(s)
}

func TestConv(t *testing.T)  {
  s := strconv.Itoa(10)
  fmt.Println(s)

  if n, e := strconv.Atoi(s); e==nil {
    fmt.Println(n, e)
  }
}
