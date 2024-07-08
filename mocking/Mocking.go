package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	count = 3
	write = "write"
	sleep = "sleep"
)

type Sleeper interface {
	Sleep()
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountDownOperations struct {
	Calls []string
}
type ConfiguarableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfiguarableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	finalWord = "Blast off!"
)

func main() {
	sleeper := &ConfiguarableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)

}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := count; i >= 1; i-- {
		fmt.Fprintf(w, "%d \n", i)
		sleeper.Sleep()
	}
	fmt.Fprintf(w, "%v", finalWord)
}
