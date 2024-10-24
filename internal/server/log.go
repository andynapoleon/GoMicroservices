package server

import (
	"fmt"
	"sync"
)

// Log object is just a slice of records
type Log struct {
	mu      sync.Mutex // mutex lock
	records []Record
}

// Create a pointer to a new Log object with an empty record slice
func NewLog() *Log {
	return &Log{}
}

// Append a record to the end of the log
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock() // this will be called when the function returns even if panic
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read a record from the log based on the specified offset
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

// A Record object that has a slice of values and an offset
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
