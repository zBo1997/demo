package main

import "testing"

var benchmarkResult int

// ============ 单个参数基准测试 ============

// 基准测试：指针版本（单参数）
func BenchmarkAddPointerSingle(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := 1
		benchmarkResult = Add(&p, 1)
	}
}

// 基准测试：普通版本（单参数）
func BenchmarkAddNormalSingle(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkResult = AddNormal(1, 1)
	}
}

// ============ 多个参数基准测试 ============

// 基准测试：指针版本（多参数 5 个）
func BenchmarkAddPointerMulti(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p := 1
		benchmarkResult = Add(&p, 1, 2, 3, 4, 5)
	}
}

// 基准测试：普通版本（多参数 5 个）
func BenchmarkAddNormalMulti(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkResult = AddNormal(1, 1, 2, 3, 4, 5)
	}
}
