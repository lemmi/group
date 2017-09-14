# group

[![GoDoc](https://godoc.org/github.com/lemmi/group?status.svg)](https://godoc.org/github.com/lemmi/group)

Group slices with a less-function or collections that implement a subset of the
`sort.Interface`.

## Example

```go
package main

import (
    "fmt"
    "sort"

    "github.com/lemmi/group"
)

func main() {
    s := sort.IntSlice([]int{2, 2, 1, 2, 1, 1})
    s.Sort()

    var groups [][]int
    var g group.Grouper
    for g.Scan(s) {
        group := s[g.L:g.R]
        groups = append(groups, group)
        // ...
    }

    fmt.Println(groups)
    // [[1 1 1] [2 2 2]]
}
```
