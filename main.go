package main

import (
	"HelloWorld/mocking"
	"os"
)

func main() {
	sleeper := &mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}
