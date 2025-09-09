package main

// rl "github.com/gen2brain/raylib-go/raylib"

type Timer struct {
	interval float32 // how often to trigger (seconds)
	elapsed  float32
}

func NewTimer(interval float32) Timer {
	return Timer{interval: interval, elapsed: 0}
}

func (t *Timer) Update() bool {
	t.elapsed += dt
	if t.elapsed >= t.interval {
		t.elapsed -= t.interval
		return true
	}
	return false
}

func (t *Timer) Reset() {
	t.elapsed = 0
}
