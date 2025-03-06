package main

import (
	"HelloWorld/mocking"
	"os"
	"time"
)

func main() {
	mock()
}

func mock() {
	/*sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)*/

	sleeper := &mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
