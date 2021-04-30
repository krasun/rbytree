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
	} else {
		fmt.Println("value for banana not found")
	}

	tree.ForEach(func(key, value []byte) {
		fmt.Printf("key = %s, value = %s\n", string(key), string(value))
	})

	// Output: 
	// banana = honey
	// key = apple, value = sweet
	// key = banana, value = honey
	// key = cinnamon, value = savoury
}
```

## Use cases 

1. When you want to use []byte as a key in the map. 
2. When you want to iterate over keys in map in sorted order.

## Limitations 

**Caution!** To guarantee that the red-black tree properties are not violated, keys are copied. 

You should clearly understand what []byte slice is and why it is dangerous to use it as a key. Go language authors do prohibit using byte slice ([]byte) as a map key for a reason. The point is that you can change the values of the key and thus violate the invariants of map: 

```go
// if it worked 
b := []byte{1}
m := make(map[[]byte]int)
m[b] = 1

b[0] = 2 // it would violate the invariants 
m[[]byte{1}] // what do you expect to receive?
```

So to make sure that this situation does not occur in the tree, the key is copied byte by byte.

## Benchmark

Regular Go map is as twice faster for put and get than red-black tree. But if you 
need to iterate over keys in sorted order, the picture is sligthly different: 

```
$ go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/krasun/rbytree
BenchmarkTreePut-8                     	     332	   3538876 ns/op	 1040040 B/op	   49902 allocs/op
BenchmarkMapPut-8                      	     500	   2399698 ns/op	 1732282 B/op	   20150 allocs/op
BenchmarkTreePutRandomized-8           	     288	   4188292 ns/op	 1040028 B/op	   49899 allocs/op
BenchmarkMapPutRandomized-8            	     626	   1879120 ns/op	  981475 B/op	   20110 allocs/op
BenchmarkMapGet-8                      	    1612	    725188 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkTreeGet-8                     	     783	   1543038 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkTreePutAndForEach-8           	     312	   3814416 ns/op	 1040039 B/op	   49902 allocs/op
BenchmarkMapPutAndIterateAfterSort-8   	     214	   5982573 ns/op	 2558689 B/op	   20174 allocs/op
PASS
ok  	github.com/krasun/rbytree	12.478s
```

## Tests

Run tests with: 

```
$ go test -cover .
ok  	github.com/krasun/rbytree	0.245s	coverage: 100.0% of statements
```

## License 

**rb**y**tree** is released under [the MIT license](LICENSE).