package leven

import (
	"fmt"
	"testing"

	"github.com/go-ego/gse"
	"github.com/vcaesar/tt"
)

var seg gse.Segmenter

func init() {
	seg.LoadDict()
}

func Test_Seq(t *testing.T) {
	p := *NewParams().FilterScore(0.2)
	s1 := "City of Seattle, 西雅图都会区"
	s2 := "The Space Nedle, 西雅图太空针"

	r := Ratio(s1, s2, p)
	tt.Equal(t, 0.34782608695652173, r)

	m := MatchNew(s1, s2)
	tt.Equal(t, 0.34782608695652173, m)

	d := DistanceNew(s1, s2)
	tt.Equal(t, 15, d)

	seq1 := seg.Cut(s1)
	seq2 := seg.Cut(s2)

	s := SeqRatio(seq1, seq2, p)
	fmt.Println("cut with ratio: ", s)
	tt.Equal(t, 0.5014285714285716, s)
}
