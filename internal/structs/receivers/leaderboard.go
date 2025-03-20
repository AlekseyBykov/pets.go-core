package receivers

import (
	"unsafe"
)

type Player struct {
	Score int
}

type Leaderboard struct {
	Players [10_000_000]Player
}

func (lb *Leaderboard) addr() uintptr {
	return uintptr(unsafe.Pointer(&lb.Players[0]))
}

func (lb Leaderboard) addrCopy() uintptr {
	return uintptr(unsafe.Pointer(&lb.Players[0]))
}
