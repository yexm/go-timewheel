package example

import (
	"fmt"
	"time"

	"go-timewheel/src"
)

var (
	DefaultTimeWheel, _ = src.NewTimeWheel(time.Second, 120)
)

func init() {
	DefaultTimeWheel.Start()
}

func ResetDefaultTimeWheel(tw *src.TimeWheel) {
	tw.Start()
	DefaultTimeWheel = tw
}

func Add(delay time.Duration, callback func()) *src.Task {
	return DefaultTimeWheel.Add(delay, callback)
}

func AddCron(delay time.Duration, callback func()) *src.Task {
	return DefaultTimeWheel.AddCron(delay, callback)
}

func Remove(task *src.Task) error {
	return DefaultTimeWheel.Remove(task)
}

func NewTimer(delay time.Duration) *src.Timer {
	return DefaultTimeWheel.NewTimer(delay)
}

func NewTicker(delay time.Duration) *src.Ticker {
	return DefaultTimeWheel.NewTicker(delay)
}

func AfterFunc(delay time.Duration, callback func()) *src.Timer {
	return DefaultTimeWheel.AfterFunc(delay, callback)
}

func After(delay time.Duration) <-chan time.Time {
	return DefaultTimeWheel.After(delay)
}

func Sleep(delay time.Duration) {
	fmt.Println(3)
	DefaultTimeWheel.Sleep(delay)
}
