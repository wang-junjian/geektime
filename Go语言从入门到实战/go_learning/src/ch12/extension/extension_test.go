package extension_test

import (
  "testing"
  "fmt"
)

type Pet struct {

}

func (p *Pet) Speak()  {
  fmt.Println("...")
}

func (p *Pet) SpeakTo(name string) {
  p.Speak()
  fmt.Println(" ", name)
}

type Dog struct {
  Pet
}

func (d *Dog) Speak()  {
  fmt.Print("Wang!")
}

func TestDog(t *testing.T)  {
  dog := new(Dog)
  dog.SpeakTo("wjj")
}
