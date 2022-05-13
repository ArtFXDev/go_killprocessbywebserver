package utils

import (
	"time"
)

type SetInterval struct {
	// private
	ticker        *time.Ticker
	tickerChannel chan bool
	isRunning     bool
}

func (si *SetInterval) Start(someFunc func(), milliseconds int, async bool) chan bool {

	// How often to fire the passed in function
	// in milliseconds
	interval := time.Duration(milliseconds) * time.Millisecond

	// Setup the ticket and the channel to signal
	// the ending of the interval
	si.ticker = time.NewTicker(interval)
	si.tickerChannel = make(chan bool)

	si.isRunning = true
	// Put the selection in a go routine
	// so that the for loop is none blocking
	go func() {
		for {
			select {
			case <-si.tickerChannel:
				si.isRunning = false
				si.ticker.Stop()
				return

			case <-si.ticker.C:
				si.isRunning = true
				if async {
					// This won't block
					go someFunc()
				} else {
					// This will block
					someFunc()
				}
			}
		}
	}()

	// We return the channel so we can pass in
	// a value to it to clear the interval
	return si.tickerChannel
}

func (si *SetInterval) Stop() {
	si.tickerChannel <- true
}

func (si *SetInterval) IsRunning() bool {
	return si.isRunning
}
