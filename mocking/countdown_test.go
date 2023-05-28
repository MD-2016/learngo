package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 5 to Go!", func(t *testing.T) {
		buff := &bytes.Buffer{}
		Countdown(buff, &SpyCountdownOp{})

		got := buff.String()
		want := `5
4
3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		printerSpy := &SpyCountdownOp{}
		Countdown(printerSpy, printerSpy)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, printerSpy.Calls) {
			t.Errorf("wanted calls %v got %v", want, printerSpy.Calls)
		}
	})
}

func TestConfigSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

type SpyCountdownOp struct {
	Calls []string
}

func (s *SpyCountdownOp) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOp) Write(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(dur time.Duration) {
	s.durationSlept = dur
}
