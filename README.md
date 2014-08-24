# go-vache

be inspired by https://metacpan.org/release/Cache-Scalar-WithExpiry

## SYNOPSIS

```go
package main

import (
	"log"
	"time"

	"github.com/hisaichi5518/vache"
)

func main() {
	vache.Set("key", "val", time.Second)
	v, ok := vache.Get("key")
	log.Println(v)  //=> 2014/08/23 19:58:22 val  <= Having Cache!
	log.Println(ok) //=> 2014/08/23 19:58:22 true <= Having Cache!

    // Wait for expire..
	time.Sleep(2 * time.Second)
	v, ok = vache.Get("key")
	if !ok {
		v = "not found"
	}
	log.Println(v)  //=> 2014/08/23 19:58:24 not found <= Deleted!
	log.Println(ok) //=> 2014/08/23 19:58:24 false     <= Deleted!
}
```

## Install

To install, use `go get`:

```bash
$ go get github.com/hisaichi5518/vache
```

## Contribution

1. Fork ([https://github.com/hisaichi5518/vache/fork](https://github.com/hisaichi5518/vache/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[hisaichi5518](https://github.com/hisaichi5518)
