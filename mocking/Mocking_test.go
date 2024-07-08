package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Prints correctly", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountDownOperations{})

		got := buffer.String()
		want := "3 \n2 \n1 \nBlast off!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleeps before prints", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("expected calls %v, got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfiguarableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}
