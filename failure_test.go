package failured_test

import (
	"testing"

	. "github.com/andy2046/failured"
)

func withDiffs(fd *Detector, samples ...int64) int64 {
	time := int64(0)
	fd.RegisterHeartbeat(0)

	for _, s := range samples {
		time += s
		fd.RegisterHeartbeat(time)
	}

	return time
}

func TestNoVariation(t *testing.T) {
	fd := New()
	lastTime := withDiffs(fd, 1, 1, 1, 1)

	data := []struct {
		expectedVal  float64
		expectedBool bool
		v            int64
	}{
		{0, false, lastTime},
		{0, false, lastTime + 1},
		{1, true, lastTime + 2},
	}

	for _, d := range data {
		p := fd.FailureProbability(d.v)
		if p != d.expectedVal {
			t.Fatalf("expected %v got %v", d.expectedVal, p)
		}
		b := fd.CheckFailure(d.v)
		if b != d.expectedBool {
			t.Fatalf("expected %v got %v", d.expectedBool, b)
		}
	}

}

func TestVariation(t *testing.T) {
	fd := New()
	lastTime := withDiffs(fd, 1010, 1023, 1012, 1032, 1016, 1020, 990, 1028)

	data := []struct {
		expectedVal  float64
		expectedBool bool
		v            int64
	}{
		{0, false, lastTime + 500},
		{0, false, lastTime + 1000},
		{0.125, false, lastTime + 1100},
		{1, true, lastTime + 2100},
	}

	for _, d := range data {
		p := fd.FailureProbability(d.v)
		if p != d.expectedVal {
			t.Fatalf("expected %v got %v", d.expectedVal, p)
		}
		b := fd.CheckFailure(d.v)
		if b != d.expectedBool {
			t.Fatalf("expected %v got %v", d.expectedBool, b)
		}
	}

}
