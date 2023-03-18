package bbloom

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBloomTrace(t *testing.T) {
	var filter = New(1000, 0.01)
	var trace = filter.DeriveTrace()

	trace.Add([]byte("foo"))
	trace.Add([]byte("bar"))

	trace.SyncTo(filter)

	assert.Equal(t, uint64(0xe), trace.Length(), "trace length should be uint64(0xe when inserting foo and bar")
	assert.Equal(t, uint64(0xe), filter.ElemNum, "filter length should be uint64(0xe when inserting foo and bar")

	var anotherFilter = New(1000, 0.01)
	trace.SyncTo(anotherFilter)

	assert.Equal(t, uint64(0xe), anotherFilter.ElemNum, "anotherFilter should have 14 elements after sync")
	assert.Equal(t, filter.JSONMarshal(), anotherFilter.JSONMarshal())
}
