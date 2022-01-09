# **rb**y**tree**

[![Build](https://github.com/krasun/rbytree/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/krasun/rbytree/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/krasun/rbytree/branch/main/graph/badge.svg?token=8NU6LR4FQD)](https://codecov.io/gh/krasun/rbytree)
[![Go Report Card](https://goreportcard.com/badge/github.com/krasun/rbytree)](https://goreportcard.com/report/github.com/krasun/rbytree)
[![GoDoc](https://godoc.org/https://godoc.org/github.com/krasun/rbytree?status.svg)](https://godoc.org/github.com/krasun/rbytree)

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

You can use an iterator: 

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

	for it := tree.Iterator(); it.HasNext(); {
		key, value := it.Next()
		fmt.Printf("key = %s, value = %s\n", string(key), string(value))
	}

	// Output: 
	// banana = honey
	// key = apple, value = sweet
	// key = banana, value = honey
	// key = cinnamon, value = savoury
}
```

An iterator is stateful. You can have multiple iterators without any impact on each other, but make sure to synchronize access to them and the tree in a concurrent environment.

Caution! `Next` panics if there is no next element. Make sure to test for the next element with `HasNext` before.

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
BenchmarkTreePut-8                     	     330	   3573752 ns/op	 1040040 B/op	   49902 allocs/op
BenchmarkMapPut-8                      	     496	   2477226 ns/op	 1732586 B/op	   20151 allocs/op
BenchmarkTreePutRandomized-8           	     260	   4394145 ns/op	 1040029 B/op	   49901 allocs/op
BenchmarkMapPutRandomized-8            	     630	   1890784 ns/op	  981565 B/op	   20111 allocs/op
BenchmarkMapGet-8                      	    1496	    768210 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkTreeGet-8                     	     723	   1604544 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkTreePutAndForEach-8           	     300	   4056864 ns/op	 1040043 B/op	   49903 allocs/op
BenchmarkMapPutAndIterateAfterSort-8   	     202	   5559646 ns/op	 2558408 B/op	   20173 allocs/op
PASS
ok  	github.com/krasun/rbytree	12.096s
```

## Tests

Run tests with: 

```
$ go test -cover .
ok  	github.com/krasun/rbytree	0.245s	coverage: 100.0% of statements
```

## Known Usages 

1. [krasun/lsmtree](https://github.com/krasun/lsmtree) - my experimental implementation of [log-structured merge-tree](https://en.wikipedia.org/wiki/Log-structured_merge-tree).

## License 

**rb**y**tree** is released under [the MIT license](LICENSE).