package bbloom

import "sync"

// BloomTrace -- BloomTrace related implementations
type BloomTrace struct {
	records []uint64
	length  uint64
	mtx     *sync.Mutex
}

func NewTrace() *BloomTrace {
	return &BloomTrace{records: make([]uint64, 0), mtx: new(sync.Mutex)}
}

func NewTraceFromRecords(records []uint64) *BloomTrace {
	return &BloomTrace{records: records, mtx: new(sync.Mutex)}
}

func (bt *BloomTrace) Length() uint64 {
	return bt.length
}

func (bt *BloomTrace) Add(trace uint64) {
	bt.records = append(bt.records, trace)
	bt.length++
}

func (bt *BloomTrace) AddTS(trace uint64) {
	bt.mtx.Lock()
	defer bt.mtx.Unlock()
	bt.Add(trace)
}

func (bt *BloomTrace) SyncTo(bloom *Bloom) {
	for _, trace := range bt.records {
		bloom.Sync(trace)
	}
}

func (bt *BloomTrace) SyncToTS(bloom *Bloom) {
	bloom.Mtx.Lock()
	defer bloom.Mtx.Unlock()
	bt.SyncTo(bloom)
}

// Sync
// Sync sets bitset from bloomTrace. This function is intended to be used
// when syncing bloom filters incrementally over the wire
func (bl *Bloom) Sync(record uint64) {
	bl.set(record)
	bl.ElemNum++
}

func (bl *Bloom) AddWithTrace(entry []byte, trace *BloomTrace) {
	l, h := bl.sipHash(entry)

	for i := uint64(0); i < bl.setLocs; i++ {
		x := (h + i*l) & bl.size
		bl.set(x)

		trace.Add(x)
		bl.ElemNum++
	}
}
