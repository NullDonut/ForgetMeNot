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
func memAt(stability float64, hoursAgo float64) Memory {
	return Memory{Stability: stability, TimeLearned: time.Now().Add(-time.Duration(hoursAgo * float64(time.Hour)))}
}
func TestReinforceTooFresh(t *testing.T) {
	m := memAt(24, 0)
	if got := memoryReinforce(&m); got != 1.0 || m.ReviewCount != 0 {
		t.Fatalf("expected no-op,got mult=%v reviews=%d", got, m.ReviewCount)
	}
}
func TestReinforceGrows(t *testing.T) {
	m := memAt(24, 48)
	if got := memoryReinforce(&m); got != 2.0 || m.Stability != 48 {
		t.Fatalf("got mult=%v stability=%v", got, m.Stability)
	}
}
func TestLapseSurvives(t *testing.T) {
	m := memAt(24, 60)
	memoryLapse(&m)
	if MemoryStrength(m) < 0.29 {
		t.Fatalf("Memory would be deleted,strength=%v", MemoryStrength(m))
	}

}
