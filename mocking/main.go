package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigSleeper) Sleep() {
	c.sleep(c.duration)
}

const endWord = "Go!"
const countdownStart = 5

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprintf(out, endWord)
}

func main() {
	sleeper := &ConfigSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
