package snowflake

import (
	"sync"
	"time"
)

var (
	workerIDBits uint64 = 10
	maxWorkerID         = -1 ^ (-1 << workerIDBits)
	sequenceBits uint64 = 12
	sequenceMask        = -1 ^ (-1 << sequenceBits)

	workerIDShift      = sequenceBits
	timeStampLeftShift = workerIDBits + sequenceBits
)

// IDWorker generates IDs
type IDWorker struct {
	mu       sync.Mutex
	workerID uint64
	sequence uint64
	twepoch  int64 // 相对于

	lastTimestamp int64
}

// NewIDWorker new
func NewIDWorker(workerID uint64) *IDWorker {
	if workerID > uint64(maxWorkerID) {
		return nil
	}
	return &IDWorker{
		workerID:      workerID,
		sequence:      0,
		twepoch:       0,
		lastTimestamp: -1,
	}
}

// NextID return next id
func (w *IDWorker) NextID() uint64 {
	timestamp := time.Now().UnixNano() / 1e6

	w.mu.Lock()
	defer w.mu.Unlock()
	if timestamp < w.lastTimestamp {
		panic("clock is moving backwards. Reject request")
	} else if timestamp == w.lastTimestamp {
		w.sequence = (w.sequence + 1) & uint64(sequenceMask)
		if w.sequence == 0 {
			// 等到下一毫秒
			for timestamp <= w.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.sequence = 0
	}
	w.lastTimestamp = timestamp
	return uint64(timestamp-w.twepoch)<<timeStampLeftShift |
		w.workerID<<workerIDShift |
		w.sequence
}
