package utils

import "sync/atomic"

type AtomicInteger struct {
	value int32
}

func NewAtomicInteger(value int32) *AtomicInteger {
	return &AtomicInteger{value: value}
}

func (ai *AtomicInteger) IncAndGet() int {
	return int(atomic.AddInt32(&ai.value, 1))
}

func (ai *AtomicInteger) Set(value int) {
	atomic.StoreInt32(&ai.value, int32(value))
}
