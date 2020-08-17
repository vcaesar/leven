// Copyright 2016 vcaesar Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leven

// Ratio return s1, s2 ratio
func Ratio(s1, s2 string, p ...Params) float64 {
	p1 := NewParams()
	if len(p) > 0 {
		p1 = &p[0]
	}

	return Similarity(s1, s2, p1)
}

// MatchNew match with new params
func MatchNew(s1, s2 string) float64 {
	return Match(s1, s2, NewParams())
}

// DistanceNew distance with new params
func DistanceNew(s1, s2 string) int {
	return Distance(s1, s2, NewParams())
}

// MatchSeq match struct
type MatchSeq struct {
	Index1, Index2 int
	Matched        bool

	Ratio    float64
	Distance int
}

// MatchMatrix match for matrix matchseq
func MatchMatrix(s1, s2 []string, p Params) (match [][]MatchSeq, f float64) {
	len1 := len(s1)
	len2 := len(s2)
	matched := make(map[int]bool)

	for i := 0; i < len1; i++ {
		var match1 []MatchSeq
		for h := 0; h < len2; h++ {
			// d := Distance(s1[i], s2[h], &p)
			r := Similarity(s1[i], s2[h], &p)
			if r > p.filterScore {
				m1 := MatchSeq{
					Index1: i,
					Index2: h,
					Ratio:  r,
					// Distance: d,
				}

				if r == 1 && i == h {
					f += r
					m1.Matched = true
					matched[h] = true
					match1 = append(match1, m1)
					break
				}

				if !matched[h] {
					f += r
					match1 = append(match1, m1)
				}
			}
		}

		if len(match1) > 0 {
			match = append(match, match1)
		}
	}

	return match, f
}

// SeqRatio return s1, s2 sequence ratio
func SeqRatio(s1, s2 []string, p Params) (f float64) {
	matched := make(map[int]bool)
	h1 := 0

	m, _ := MatchMatrix(s1, s2, p)
	for i := 0; i < len(m); i++ {
		if len(m[i]) == 1 {
			f += m[i][0].Ratio
			h1++
			matched[m[i][0].Index2] = true
		} else {
			maxVal := m[i][0].Ratio
			maxIndex := 0
			for h := 0; h < len(m[i]); h++ {
				if m[i][h].Ratio == 1.0 {
					maxIndex = h
					matched[m[i][maxIndex].Index2] = false
					break
				}

				if maxVal < m[i][h].Ratio {
					maxVal = m[i][h].Ratio
					maxIndex = h
				}

				if maxVal == m[i][h].Ratio && matched[m[i][maxIndex].Index2] {
					maxIndex = h
				}
			}

			if !matched[m[i][maxIndex].Index2] {
				f += m[i][maxIndex].Ratio
				h1++
				matched[m[i][maxIndex].Index2] = true
			}
		}
	}

	l := len(s1) + len(s2) - h1*2
	f = f/float64(len(s1)+len(s2))*2 - float64(l)*0.01

	return
}
