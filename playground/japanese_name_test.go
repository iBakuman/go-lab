package playground

import (
	"testing"

	"github.com/mattn/go-gimei"
)

func TestGetName(t *testing.T) {
	name := gimei.NewName()
	t.Logf(name.First.Kanji())
	t.Logf(name.First.Katakana())
	t.Logf(name.Last.Kanji())
	t.Logf(name.Last.Katakana())
}
