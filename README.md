# rbytree

[![Build Status](https://travis-ci.com/krasun/rbytree.svg?branch=main)](https://travis-ci.com/krasun/rbytree)
[![codecov](https://codecov.io/gh/krasun/rbytree/branch/main/graph/badge.svg?token=8NU6LR4FQD)](https://codecov.io/gh/krasun/rbytree)
[![Go Report Card](https://goreportcard.com/badge/github.com/krasun/rbytree)](https://goreportcard.com/report/github.com/krasun/rbytree)
[![GoDoc](https://godoc.org/https://godoc.org/github.com/krasun/rbytree?status.svg)](https://godoc.org/github.com/krasun/rbytree)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkrasun%2Frbytree.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkrasun%2Frbytree?ref=badge_shield)

A [red-black tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree) implementation for Go with byte-slice keys and values. 

## Installation 

To install, run:
```
go get github.com/krasun/bytree
```

## Quickstart

Feel free to play: 
```

```

## Limitations 

...

## Benchmark

```
$ go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/krasun/rbtree
BenchmarkTreePut-8                     	     355	   3251161 ns/op	 1039232 B/op	   39902 allocs/op
BenchmarkTreePutRandomized-8           	     310	   3896154 ns/op	 1039221 B/op	   39901 allocs/op
BenchmarkTreePutAndForEach-8           	     364	   3249646 ns/op	 1039230 B/op	   39902 allocs/op
BenchmarkTreeGet-8                     	     774	   1677112 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkMapPut-8                      	     480	   2638418 ns/op	 1732639 B/op	   20151 allocs/op
BenchmarkMapPutRandomized-8            	     582	   2145951 ns/op	  981695 B/op	   20111 allocs/op
BenchmarkMapGet-8                      	    1552	    806005 ns/op	   38880 B/op	    9900 allocs/op
BenchmarkMapPutAndIterateAfterSort-8   	     180	   6447948 ns/op	 2557653 B/op	   20171 allocs/op
PASS
ok  	github.com/krasun/rbtree	13.273s
```

## License 

rbytree is released under [the MIT license](LICENSE).