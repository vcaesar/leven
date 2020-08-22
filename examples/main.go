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
