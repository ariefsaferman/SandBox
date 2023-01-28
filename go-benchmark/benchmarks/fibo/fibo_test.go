package fibo

import (
	"testing"
)

func TestSequentialFibonacci(t *testing.T) {
	data := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{42, 267914296},
	}

	for _, d := range data {
		if got := SequentialFibonacci(d.n); got != d.want {
			t.Errorf("got: %d, want: %d", got, d.want)
		}
	}
}

func TestRecursiveFibonacci(t *testing.T) {
	data := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{10, 55},
		{42, 267914296},
	}

	for _, d := range data {
		if got := RecursiveFibonacci(d.n); got != d.want {
			t.Errorf("got: %d, want: %d", got, d.want)
		}
	}
}

func BenchmarkTestRecursiveFibonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibonacci(10)
	}
}

func BenchmarkTestRecursiveFibonacci_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibonacci(20)
	}
}

func BenchmarkTestRecursiveFibonacci_40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveFibonacci(40)
	}
}

func BenchmarkSequentialFibonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(10)
	}
}

func BenchmarkSequentialFibonacci_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(20)
	}
}

func BenchmarkSequentialFibonacci_40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(40)
	}
}

func BenchmarkSequentialFibonacci_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SequentialFibonacci(1000)
	}
}
