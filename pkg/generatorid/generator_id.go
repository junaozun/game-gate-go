package generatorid

import (
	"sync/atomic"
)

type IDGenerator interface {
	NextID() uint64
}

// IncIDGenerator 基本实现
type IncIDGenerator uint64

func (g *IncIDGenerator) NextID() uint64 {
	return atomic.AddUint64((*uint64)(g), 1)
}
