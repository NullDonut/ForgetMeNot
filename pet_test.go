package main

import (
	"math"
	"testing"
	"time"
)

func TestStrengthAtOneStability(t *testing.T) {
	m := Memory{
		Stability:   24,
		TimeLearned: time.Now().Add(-24 * time.Hour),
	}
	if got := MemoryStrength(m); math.Abs(got-math.Exp(-1)) > 0.01 {
		t.Fatalf("got %v, want ~0.37", got)
	}
}
