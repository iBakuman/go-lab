package unicode

import (
	"strings"
	"testing"
	"unicode"
)

func TestPrintAllHiragana(t *testing.T) {
	total := 0
	strBuilder := strings.Builder{}
	for _, r16 := range unicode.Hiragana.R16 {
		for r := r16.Lo; r <= r16.Hi; r += r16.Stride {
			strBuilder.WriteRune(rune(r))
			total++
		}
	}
	for _, r32 := range unicode.Hiragana.R32 {
		for r := r32.Lo; r <= r32.Hi; r += r32.Stride {
			strBuilder.WriteRune(rune(r))
			total++
		}
	}
	t.Log(strBuilder.String())
	t.Logf("total: %d", total)
}

func TestPrintAllKataKana(t *testing.T) {
	total := 0
	strBuilder := strings.Builder{}
	for _, r16 := range unicode.Katakana.R16 {
		for r := r16.Lo; r <= r16.Hi; r += r16.Stride {
			strBuilder.WriteRune(rune(r))
			total++
		}
	}
	for _, r32 := range unicode.Katakana.R32 {
		for r := r32.Lo; r <= r32.Hi; r += r32.Stride {
			strBuilder.WriteRune(rune(r))
			total++
		}
	}
	t.Log(strBuilder.String())
	t.Logf("total: %d", total)
}

func TestPrintHalfWidthKatakana(t *testing.T) {
	total := 0
	strBuilder := strings.Builder{}
	var start = '\uFF61'
	for start <= '\uFF9F' {
		strBuilder.WriteRune(start)
		start++
		total++
	}
	t.Log(strBuilder.String())
	t.Logf("total: %d", total)
}

func TestName(t *testing.T) {
	name := "ゆう子"
	for _, ch := range name {
		if unicode.In(ch, unicode.Katakana) {
			t.Logf("%c is katakana", ch)
		} else if unicode.In(ch, unicode.Hiragana) {
			t.Logf("%c is hiragana", ch)
		} else if unicode.In(ch, unicode.Han) {
			t.Logf("%c is kanji", ch)
		} else {
			t.Logf("%c is not japanese", ch)
		}
	}
}
