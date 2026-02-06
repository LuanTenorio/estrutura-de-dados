package dynamic_test

import (
	"testing"

	"github.com/LuanTenorio/estrutura-de-dados/dynamic"
)

const benchN = 30

func BenchmarkComputingFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamic.ComputingFib(benchN)
	}
}

func BenchmarkCachyFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamic.CachyFib(benchN)
	}
}

func BenchmarkDPFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dynamic.DPFib(benchN)
	}
}
