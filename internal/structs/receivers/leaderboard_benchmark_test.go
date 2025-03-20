package receivers

import "testing"

func BenchmarkLeaderboardPointerReceiver(b *testing.B) {
	lb := Leaderboard{}
	var result uintptr // blackhole

	for i := 0; i < b.N; i++ {
		result = lb.addr()
	}

	_ = result
}

func BenchmarkLeaderboardValueReceiver(b *testing.B) {
	lb := Leaderboard{}
	var result uintptr // blackhole

	for i := 0; i < b.N; i++ {
		result = lb.addrCopy()
	}

	_ = result
}
