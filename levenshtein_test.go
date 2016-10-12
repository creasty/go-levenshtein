package levenshtein

import (
	"testing"
)

func TestDistance(t *testing.T) {
	samples := map[[2]string]int{
		[2]string{"", ""}:                        0,
		[2]string{"", "a"}:                       1,
		[2]string{"a", ""}:                       1,
		[2]string{"a", "a"}:                      0,
		[2]string{"a", "b"}:                      1,
		[2]string{"kitten", "sitting"}:           3,
		[2]string{"example", "samples"}:          3,
		[2]string{"creasty", "Creasty"}:          1,
		[2]string{"levenshtein", "frankenstein"}: 6,
		[2]string{"distance", "difference"}:      5,
		[2]string{"あいうえお", "aいうえo"}:              2,
	}

	for pair, expect := range samples {
		actual := Distance(pair[0], pair[1])
		if actual != expect {
			t.Errorf("Expected the distance of `%s` and `%s` to be %d, but was %d", pair[0], pair[1], expect, actual)
		}
	}
}

func BenchmarkDistance(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Distance("あいうえお levenshtein", "aiueo frankenstein")
	}
}

func TestLcsLen(t *testing.T) {
	samples := map[[2]string]int{
		[2]string{"", ""}:                        0,
		[2]string{"", "a"}:                       0,
		[2]string{"a", ""}:                       0,
		[2]string{"a", "a"}:                      1,
		[2]string{"a", "b"}:                      0,
		[2]string{"kitten", "sitting"}:           4,
		[2]string{"example", "samples"}:          5,
		[2]string{"creasty", "Creasty"}:          6,
		[2]string{"levenshtein", "frankenstein"}: 7,
		[2]string{"distance", "difference"}:      5,
		[2]string{"あいうえお", "aいうえo"}:              3,
	}

	for pair, expect := range samples {
		actual := LcsLen(pair[0], pair[1])
		if actual != expect {
			t.Errorf("Expected the distance of `%s` and `%s` to be %d, but was %d", pair[0], pair[1], expect, actual)
		}
	}
}
