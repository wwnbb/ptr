Well, i know that its weird to create packages for a single line function.

But its way weirder to have a single line function in every project you have.


```go

import "github.com/wwnbb/ptr"

type Stupid struct {
  p *int
}

func NewStupid() Stupid {
  i := 0
  return Stupid{&i}
}

func NewSmart() Stupid {
  return Stupid(ptr(0))
}
```
