package main

import "testing"

func BenchmarkLen(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		_ = len(s)
	}
}

func BenchmarkEqual(b *testing.B) {
	s := ""
	for i := 0; i < b.N; i++ {
		_ = s == ""
	}
}
