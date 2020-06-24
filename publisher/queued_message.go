package main

import (
	"fmt"
	"time"
)

const IntervalPeriod = 3 * time.Second


type jobTicker struct {
	t *time.Timer
}

func getNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(IntervalPeriod)
	}
	return nextTick.Sub(time.Now())
}

func NewJobTicker() jobTicker {
	fmt.Println("new tick here")
	return jobTicker{time.NewTimer(getNextTickDuration())}
}

func (jt jobTicker) updateJobTicker() {
	fmt.Println("next tick here")
	jt.t.Reset(getNextTickDuration())
}