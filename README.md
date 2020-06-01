# cmap

[![DOI](https://zenodo.org/badge/159430285.svg)](https://zenodo.org/badge/latestdoi/159430285)
[![GoDoc](https://godoc.org/github.com/kchristidis/cmap?status.svg)](https://godoc.org/github.com/kchristidis/cmap)
[![Build Status](https://travis-ci.org/kchristidis/cmap.svg?branch=master)](https://travis-ci.org/kchristidis/cmap)

A thread-safe capacity-constrained hash table that evicts key/value pairs
according to the LRU (least recently used) policy. It is technically a thin
wrapper around Go's built-in map type.

## Usage

```go
package main

import (
    cmap "github.com/kchristidis/cmap
)

func main() {
    cm, err := cmap.New(2) // will hold up to 2 key/value pairs
    if err != nil {
        log.Fatal(err)
    }

    cm.Put("fooKey", "fooVal")
    v, ok := cm.Get("fooKey") // retrieve the value
    fmt.Println(ok)
    fmt.Println(v)

    cm.Put("barKey", "barVal") // retrieve value as above
    cm.Put("bazKey", "bazVal") // at this point "fooKey" is evicted

    _, ok = cm.Get("fooKey")
    fmt.Println(ok) // expected output: false
}
```

You may also want to consult the package documentation in
[GoDoc](http://godoc.org/github.com/kchristidis/cmap).

## Contributing

Contributions are welcome. Fork this library and submit a pull request.
