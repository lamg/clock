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
