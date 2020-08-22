# A Go package for calculating the Levenshtein distance between two strings

[![Build Status](https://github.com/vcaesar/leven/workflows/Go/badge.svg)](https://github.com/vcaesar/leven/commits/master)
[![CircleCI Status](https://circleci.com/gh/vcaesar/leven.svg?style=shield)](https://circleci.com/gh/vcaesar/leven)
[![codecov](https://codecov.io/gh/vcaesar/leven/branch/master/graph/badge.svg)](https://codecov.io/gh/vcaesar/leven)
[![Build Status](https://travis-ci.org/vcaesar/leven.svg?branch=master&style=flat)](https://travis-ci.org/vcaesar/leven)
<!-- [![Coverage Status](https://coveralls.io/repos/github/vcaesar/leven/badge.svg?style=flat)](https://coveralls.io/github/vcaesar/leven) -->
[![Go Report Card](https://goreportcard.com/badge/github.com/vcaesar/leven?style=flat)](https://goreportcard.com/report/github.com/vcaesar/leven)
[![Release](https://img.shields.io/github/release/vcaesar/leven.svg?style=flat)](https://github.com/vcaesar/leven/releases/latest)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/vcaesar/leven) 

This package implements distance and similarity metrics for strings, based on the Levenshtein measure, in [Go](http://golang.org).

## Use
```Go
package main

import (
	"fmt"

	"github.com/go-ego/gse"
	"github.com/vcaesar/leven"
)

var seg gse.Segmenter

func init() {
	seg.LoadDict()
}

func main() {
	p := *leven.NewParams().FilterScore(0.6)

	s1 := "City of Seattle, 西雅图都会区"
	s2 := "The Space Nedle, 西雅图太空针, 都会区"

	r := leven.Ratio(s1, s2, p)
	fmt.Println(r)

	m := leven.MatchNew(s1, s2)
	fmt.Println(m)

	d := leven.DistanceNew(s1, s2)
	fmt.Println(d)

	seq1 := seg.Cut(s1)
	seq2 := seg.Cut(s2)

	s := leven.SeqRatio(seq1, seq2, p)
	fmt.Println("cut with ratio: ", s)
}
```

## Overview

The Levenshtein `Distance` between two strings is the minimum total cost of edits that would convert the first string into the second. The allowed edit operations are insertions, deletions, and substitutions, all at character (one UTF-8 code point) level. Each operation has a default cost of 1, but each can be assigned its own cost equal to or greater than 0.

A `Distance` of 0 means the two strings are identical, and the higher the value the more different the strings. Since in practice we are interested in finding if the two strings are "close enough", it often does not make sense to continue the calculation once the result is mathematically guaranteed to exceed a desired threshold. Providing this value to the `Distance` function allows it to take a shortcut and return a lower bound instead of an exact cost when the threshold is exceeded.

The `Similarity` function calculates the distance, then converts it into a normalized metric within the range 0..1, with 1 meaning the strings are identical, and 0 that they have nothing in common. A minimum similarity threshold can be provided to speed up the calculation of the metric for strings that are far too dissimilar for the purpose at hand. All values under this threshold are rounded down to 0.

The `Match` function provides a similarity metric, with the same range and meaning as `Similarity`, but with a bonus for string pairs that share a common prefix and have a similarity above a "bonus threshold". It uses the same method as proposed by Winkler for the Jaro distance, and the reasoning behind it is that these string pairs are very likely spelling variations or errors, and they are more closely linked than the edit distance alone would suggest.

The underlying `Calculate` function is also exported, to allow the building of other derivative metrics, if needed.

## Installation

```
go get github.com/vcaesar/leven
```

## License

Package levenshtein is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
