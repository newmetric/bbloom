package bbloom

import (
	"sync"
)

// BloomTrace -- BloomTrace related implementations
type BloomTrace struct {
	records   []uint64
	bfSize    uint64
	bfShift   uint64
	bfSetLocs uint64
	mtx       *sync.Mutex
}

func (bl *Bloom) DeriveTrace() *BloomTrace {
	return &BloomTrace{
		records:   make([]uint64, 0),
		bfSize:    bl.size,
		bfShift:   bl.shift,
		bfSetLocs: bl.setLocs,
		mtx:       new(sync.Mutex),
	}
}

func (bl *Bloom) setFromTrace(trace uint64) {
	bl.set(trace)
	bl.ElemNum++
}

// Length returns the number of entries in bloomTrace
// note that entries are not the number of added entries,
// but rather the number of trace set
func (bt *BloomTrace) Length() uint64 {
	return uint64(len(bt.records))
}

// Add adds entry to bloomTrace (akin to add() in original bloom filter)
func (bt *BloomTrace) Add(entry []byte) {
	l, h := SipHash(entry, bt.bfShift)
	for i := uint64(0); i < bt.bfSetLocs; i++ {
		bt.Set((h + i*l) & bt.bfSize)
	}
}

// AddTS is a thread safe version of Add
func (bt *BloomTrace) AddTS(entry []byte) {
	bt.mtx.Lock()
	defer bt.mtx.Unlock()
	bt.Add(entry)
}

// Set sets bitset from bloomTrace (akin to set() in original bloom filter)
func (bt *BloomTrace) Set(trace uint64) {
	bt.records = append(bt.records, trace)
}

func (bt *BloomTrace) SyncTo(bloom *Bloom) {
	for _, trace := range bt.records {
		bloom.setFromTrace(trace)
	}
}

func (bt *BloomTrace) SyncToTS(bloom *Bloom) {
	bloom.Mtx.Lock()
	defer bloom.Mtx.Unlock()
	bt.SyncTo(bloom)
}
