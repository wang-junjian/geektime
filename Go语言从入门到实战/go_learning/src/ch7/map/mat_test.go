package map_test

import (
  "testing"
  "fmt"
)

func TestInitMap(t *testing.T)  {
  m1 := map[int]int{1:1, 2:4, 3:9}
  fmt.Println(m1, len(m1))

  m2 := map[int]int{}
  m2[4] = 16
  fmt.Println(m2)

  m3 := make(map[string]int, 10)
  fmt.Println(m3)
}

func TestAccessNotExistingKey(t *testing.T)  {
  m :=map[int]int{}
  fmt.Println(m)
  fmt.Println(m[2])
  fmt.Println(m)

  v, ok := m[2]
  fmt.Println(v, ok, m)

  m[2] = 22
  v, ok = m[2]
  fmt.Println(v, ok, m)
}

func TestTravelMap(t *testing.T)  {
  m := map[string]int{"one":1, "two":2}
  for k, v := range m {
    fmt.Println(k, v)
  }
}

func TestMapValueWithFunc(t *testing.T)  {
  m := map[int]func (op int)int{}
  m[1] = func(op int)int {return op}
  m[2] = func(op int)int {return op*op}
  m[3] = func(op int)int {return op*op*op}

  fmt.Println(m)

  n := 2
  for k, v := range m {
    fmt.Println(k, v(n))
  }
}

func TestSet(t *testing.T)  {
  s := map[int]bool{}
  fmt.Println(s)

  s[1] = true
  fmt.Println(s[1], s[2])
  fmt.Println(s)
}
