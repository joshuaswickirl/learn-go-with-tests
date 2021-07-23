package mocking

import (
	"fmt"
	"io"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprintln(w, finalWord)
}

type ConfigurableSleeper struct {
	Duration time.Duration
	Sleeper  func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.Sleeper(c.Duration)
}
