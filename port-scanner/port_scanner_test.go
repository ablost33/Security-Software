package port_scanner

import "testing"

func BenchmarkInitialScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InitialScan("localhost")
	}
}

func BenchmarkWideScan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WideScan("localhost")
	}
}
