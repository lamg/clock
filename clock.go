package clock

import "time"

// Clock is an abstraction of time.Now
// useful for switching between the OS clock
// and a test clock
type Clock interface {
	Now() time.Time
}

// OSClock uses the OS clock
type OSClock struct {
}

// Now returns time.Now()
func (c *OSClock) Now() (t time.Time) {
	t = time.Now()
	return
}

// TClock is a test clock
type TClock struct {
	Intv time.Duration
	Time time.Time
}

// Now returns the current time and increases it
// by Intv
func (c *TClock) Now() (t time.Time) {
	c.Time, t = c.Time.Add(c.Intv), c.Time
	return
}

// TLClock is a clock that uses a circular
// list of time.Time
type TLClock struct {
	ls []time.Time
	n  int
}

// NewTLClock creates a new TLClock using a non-empty
// slice
func NewTLClock(ls []time.Time) (c *TLClock) {
	if len(ls) > 0 {
		c = &TLClock{ls, 0}
	}
	return
}

// Now returns the current time
func (c *TLClock) Now() (t time.Time) {
	t = c.ls[c.n]
	if c.n == len(c.ls) {
		c.n = 0
	} else {
		c.n = c.n + 1
	}
	return
}
