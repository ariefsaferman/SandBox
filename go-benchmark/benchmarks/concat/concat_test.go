package concat

import (
	"testing"
)

func TestConcatkBytes(t *testing.T) {
	data := []struct {
		n    []string
		want string
	}{
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"a", "b", "c", "d"}, "abcd"},
		{[]string{"digi", "talent"}, "digitalent"},
	}

	for _, d := range data {
		if got := ConcatBytesBuffer(d.n...); got != d.want {
			t.Errorf("got: %s, want: %s", got, d.want)
		}
	}
}

func TestConcatkStrings(t *testing.T) {
	data := []struct {
		n    []string
		want string
	}{
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"a", "b", "c", "d"}, "abcd"},
		{[]string{"digi", "talent"}, "digitalent"},
	}

	for _, d := range data {
		if got := ConcatStrings(d.n...); got != d.want {
			t.Errorf("got: %s, want: %s", got, d.want)
		}
	}
}

func TestConcatFmt(t *testing.T) {
	data := []struct {
		n    []string
		want string
	}{
		{[]string{"a", "b", "c"}, "abc"},
		{[]string{"a", "b", "c", "d"}, "abcd"},
		{[]string{"digi", "talent"}, "digitalent"},
	}

	for _, d := range data {
		if got := ConcatStrings(d.n...); got != d.want {
			t.Errorf("got: %s, want: %s", got, d.want)
		}
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConcatFmtSprintf("a", "b")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConcatBytesBuffer("a", "b")
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ConcatStrings("a", "b")
	}
}
