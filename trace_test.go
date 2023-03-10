package bbloom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBloomTrace(t *testing.T) {
	var trace = NewTrace()
	var filter = New(1000, 0.01)

	filter.AddWithTrace([]byte("foo"), trace)
	filter.AddWithTrace([]byte("bar"), trace)

	assert.Equal(t, uint64(0xe), trace.Length(), "trace length should be 0xe when inserting foo and bar")

	var anotherFilter = New(1000, 0.01)
	trace.SyncTo(anotherFilter)

	assert.Equal(t, uint64(0xe), anotherFilter.ElemNum, "anotherFilter should have 0xe elements after sync")

	var anotherFilter2 = New(1000, 0.01)
	NewTraceFromRecords(trace.records).SyncToTS(anotherFilter2)
	assert.Equal(t, uint64(0xe), anotherFilter.ElemNum, "anotherFilter2 should have 0xe elements after sync")

}
