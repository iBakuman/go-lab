package playground

import (
	"strings"
	"testing"

	"github.com/arbovm/levenshtein"
)

func lcsLength(a, b string) int {
	lengths := make([]int, len(a)*len(b))
	greatestLength := 0
	for i, x := range a {
		for j, y := range b {
			if x == y {
				curr := 1
				if i != 0 && j != 0 {
					curr = lengths[(i-1)*len(b)+j-1] + 1
				}

				if curr > greatestLength {
					greatestLength = curr
				}
				lengths[i*len(b)+j] = curr
			}
		}
	}
	return greatestLength
}

func TestLevenshtein(t *testing.T) {
	identifier, password := "justin01+theplant.jp", "justin11"
	compIdentifier, compPassword := strings.ToLower(identifier), strings.ToLower(password)
	dist := levenshtein.Distance(compIdentifier, compPassword)
	lcs := float32(lcsLength(compIdentifier, compPassword)) / float32(len(compPassword))
	t.Logf("dist: %d, lcs: %f", dist, lcs)
	if dist < 5 || lcs > 0.5 {
		t.Fatal("expected distance to be greater than 5 and lcs to be less than 0.5")
	}
}
