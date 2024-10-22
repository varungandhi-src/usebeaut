package usebeaut

import "github.com/sourcegraph/beaut"

func F(a beaut.UTF8String) bool {
  r1, _ := a.RuneAt(0)
  // r2, _ := a.RuneAt(1)
  return r1 == 'a'
}

func G(a beaut.UTF8String, x bool) bool {
  if x {
    return x
  }
  return F(a)
}
