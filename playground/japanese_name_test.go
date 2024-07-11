package playground

import (
	"testing"

	"github.com/gojp/kana"
	"github.com/mattn/go-gimei"
)

func TestGetName(t *testing.T) {
	name := gimei.NewName()
	t.Logf(name.Kanji())
	t.Logf(name.First.Kanji())
	t.Logf(name.First.Katakana())
	t.Logf(name.Last.Kanji())
	t.Logf(name.Last.Katakana())
}

func TestIsKatakana(t *testing.T) {
	kana.IsKatakana()
}
