// Package failured implements the Adaptive Accrual Failure Detector
/*

This package implements [A New Adaptive Accrual Failure Detector for Dependable Distributed Systems](https://dl.acm.org/citation.cfm?id=1244129)

To use the failure detector, you need a heartbeat loop that will
call RegisterHeartbeat() at regular intervals.  At any point, you can call FailureProbability() which will
report the suspicion level since the last time RegisterHeartbeat() was called.

*/
package failured

import (
	"sync"
	"time"
	"unsafe"

	"golang.org/x/sys/cpu"
)

const (
	cacheLinePadSize = unsafe.Sizeof(cpu.CacheLinePad{})
	uint8Size        = unsafe.Sizeof(uint8(0))
	int64Size        = unsafe.Sizeof(int64(0))
)

type (
	// Detector is a failure detector.
	Detector struct {
		_         [cacheLinePadSize]byte
		nextIndex uint64
		_         [cacheLinePadSize - int64Size%cacheLinePadSize]byte

		inited uint8
		_      [cacheLinePadSize - uint8Size%cacheLinePadSize]byte

		freshness int64
		mask      uint64
		factor    float64
		threshold float64
		samples   []float64
		mu        sync.RWMutex
	}

	// Config used to init Detector.
	Config struct {
		// max size of inter-arrival time list
		WindowSize uint64
		// a scaling factor
		Factor float64

		FailureThreshold float64
	}

	// Option applies config to Detector Config.
	Option = func(*Config) error
)

var (
	// DefaultConfig is the default Detector Config.
	DefaultConfig = Config{
		WindowSize:       1000,
		Factor:           0.9,
		FailureThreshold: 0.5,
	}
)

// New returns a new failure detector.
func New(options ...Option) *Detector {
	cfg := DefaultConfig
	setOption(&cfg, options...)

	n := nextPowerOf2(cfg.WindowSize)

	return &Detector{
		mask:      n - 1,
		factor:    cfg.Factor,
		threshold: cfg.FailureThreshold,
		samples:   make([]float64, n),
	}
}

// RegisterHeartbeat registers a heartbeat at time `now`.
// `now` is the time in Milliseconds at which the heartbeat was received,
// default to time.Now() in Milliseconds.
func (d *Detector) RegisterHeartbeat(now ...int64) {
	d.mu.Lock()
	defer d.mu.Unlock()

	var tt int64
	if len(now) == 0 {
		tt = time.Now().Sub(time.Unix(0, 0)).Milliseconds()
	} else {
		tt = now[0]
	}

	if d.inited == 0 {
		d.inited = 1
		d.freshness = tt
		return
	}

	d.samples[d.nextIndex&d.mask] = float64(tt - d.freshness)
	d.nextIndex++
	d.freshness = tt
}

// FailureProbability calculates the suspicion level at time `now`
// that the remote end has failed.
// `now` is the current time in Milliseconds,
// default to time.Now() in Milliseconds.
func (d *Detector) FailureProbability(now ...int64) float64 {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.inited == 0 {
		return 0
	}

	var tt int64
	if len(now) == 0 {
		tt = time.Now().Sub(time.Unix(0, 0)).Milliseconds()
	} else {
		tt = now[0]
	}

	var (
		t     = float64(tt-d.freshness) * d.factor
		count float64
		total = min(d.nextIndex, d.mask+1)
	)

	for i := uint64(0); i < total; i++ {
		if d.samples[i] <= t {
			count++
		}
	}

	return count / float64(max(1, total))
}

// CheckFailure returns true if `FailureProbability` is
// equal to or higher than `FailureThreshold`.
// `now` is the current time in Microseconds.
func (d *Detector) CheckFailure(now ...int64) bool {
	return d.FailureProbability(now...) >= d.threshold
}

func setOption(p *Config, options ...func(*Config) error) error {
	for _, opt := range options {
		if err := opt(p); err != nil {
			return err
		}
	}
	return nil
}

func min(a, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func nextPowerOf2(v uint64) uint64 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v |= v >> 32
	return v + 1
}
