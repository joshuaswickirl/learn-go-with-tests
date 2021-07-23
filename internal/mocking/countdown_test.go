package mocking_test

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/joshuaswickirl/learn-go-with-tests/internal/mocking"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buf := &bytes.Buffer{}
		mocking.Countdown(buf, &CountdownSpy{})

		got := buf.String()
		want := "3\n2\n1\nGo!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("sleep before each write", func(t *testing.T) {
		countdownSpy := &CountdownSpy{}

		mocking.Countdown(countdownSpy, countdownSpy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, countdownSpy.Calls) {
			t.Errorf("wanted %v calls, got %v", want, countdownSpy.Calls)
		}
	})
}

type CountdownSpy struct {
	Calls []string
}

func (s *CountdownSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := mocking.ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v, slept for %v", sleepTime, spyTime.durationSlept)
	}
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
