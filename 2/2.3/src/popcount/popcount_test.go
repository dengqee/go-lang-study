package popcount

import "testing"

// func BenchmarkPopCount(b *testing.B) {
//     for i :=0; i<b.N;i++ {
//         PopCount(66666)
//     }
// }

func BenchmarkPopCount1(b *testing.B) {
    for i :=0; i<b.N;i++ {
        PopCount1(66666)
    }
}
