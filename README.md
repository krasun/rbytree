# **rb**y**tree**

[![Build Status](https://travis-ci.com/krasun/rbytree.svg?branch=main)](https://travis-ci.com/krasun/rbytree)
[![codecov](https://codecov.io/gh/krasun/rbytree/branch/main/graph/badge.svg?token=8NU6LR4FQD)](https://codecov.io/gh/krasun/rbytree)
[![Go Report Card](https://goreportcard.com/badge/github.com/krasun/rbytree)](https://goreportcard.com/report/github.com/krasun/rbytree)
[![GoDoc](https://godoc.org/https://godoc.org/github.com/krasun/rbytree?status.svg)](https://godoc.org/github.com/krasun/rbytree)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkrasun%2Frbytree.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkrasun%2Frbytree?ref=badge_shield)

A [red-black tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) implementation for Go with byte-slice keys and values. 

## Installation 

To install, run:
```
go get github.com/krasun/rbytree
```

## Quickstart

Feel free to play: 

```go
package main

import (
	"fmt"

	"github.com/krasun/rbytree"
)

func main() {
	tree := rbytree.New()

	tree.Put([]byte("apple"), []byte("sweet"))
	tree.Put([]byte("banana"), []byte("honey"))
	tree.Put([]byte("cinnamon"), []byte("savoury"))

	banana, ok := tree.Get([]byte("banana"))
	if ok {
		fmt.Printf("banana = %s\n", string(banana))
		// banana = honey
	} else {
		fmt.Println("value for banana not found")
	}

	tree.ForEach(func(key, value []byte) {
		fmt.Printf("key = %s, value = %s\n", string(key), string(value))
	})
	// key = apple, value = sweet
	// key = banana, value = honey
	// key = cinnamon, value = savoury
}
```

## Use cases 

1. When you want to use []byte as a key in the map. 
2. When you want to iterate over keys in map in sorted order.

## Limitations 

**Caution!** You can only use []byte slice keys in tree if you 100% can guarantee that the underlying array is not changed.

You should clearly understand what []byte slice is and why it is dangerous to use it as a key. Go language authors do prohibit using byte slice ([]byte) as a map key for a reason. The point is that you can change the values of the key and thus violate the invariants of map: 

```go
// if it worked 
b := []byte{1}
m := make(map[[]byte]int)
m[b] = 1

b[0] = 2 // it would violate the invariants 
m[[]byte{1}] // what do you expect to receive?
```

## Benchmark

Regular Go map is as twice faster for put and get than red-black tree. But if you 
need to iterate over keys in sorted order, the picture is sligthly different: 

```
$ go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/krasun/rbytree
BenchmarkTreePut-8                     	     350	   3349885 ns/op	 1039231 B/op	   39902 allocs/op
BenchmarkMapPut-8                      	     487	   2355075 ns/op	 1732415 B/op	   20151 allocs/op
BenchmarkTreePutRandomized-8           	     298	   3977052 ns/op	 1039221 B/op	   39901 allocs/op
BenchmarkMapPutRandomized-8            	     632	   1863399 ns/op	  981676 B/op	   20110 allocs/op
BenchmarkTreePutAndForEach-8           	     355	   3343216 ns/op	 1039230 B/op	   39902 allocs/op
BenchmarkMapGet-8                      	    1593	    736150 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkTreeGet-8                     	     768	   1519940 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkMapPutAndIterateAfterSort-8   	     208	   5720403 ns/op	 2558701 B/op	   20174 allocs/op
PASS
ok  	github.com/krasun/rbytree	12.393s
```

## License 

**rb**y**tree** is released under [the MIT license](LICENSE).