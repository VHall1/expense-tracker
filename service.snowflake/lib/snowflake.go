package lib

import (
	"fmt"
	"sync"
	"time"
)

const (
	epoch          int64 = 1732845114
	machineIDBits  uint  = 10 // Bits for machine ID
	sequenceBits   uint  = 12 // Bits for sequence
	machineIDShift uint  = sequenceBits
	timestampShift uint  = sequenceBits + machineIDBits
	maxMachineID   int64 = -1 ^ (-1 << machineIDBits)
	maxSequence    int64 = -1 ^ (-1 << sequenceBits)
)

type Snowflake struct {
	mu        sync.Mutex
	machineID int64
	lastTime  int64
	sequence  int64
}

func NewSnowflake(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, fmt.Errorf("machine ID must be between 0 and %d", maxMachineID)
	}

	return &Snowflake{
		machineID: machineID,
		lastTime:  -1,
		sequence:  0,
	}, nil
}

func (s *Snowflake) GenerateID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixMilli()
	if now == s.lastTime {
		// Increment sequence if in the same millisecond
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// Sequence overflow, wait for the next millisecond
			for now <= s.lastTime {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		// Reset sequence for a new millisecond
		s.sequence = 0
	}
	s.lastTime = now

	// Generate the ID
	id := ((now - epoch) << timestampShift) |
		(s.machineID << machineIDShift) |
		s.sequence

	return id
}
