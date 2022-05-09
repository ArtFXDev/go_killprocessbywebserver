package models

import (
	"log"
	"time"
)

type AutoMode struct {
	// private
	ticker        *time.Ticker
	tickerChannel chan bool

	// public
	Name     string
	Interval string
}
type testFunc func(chan *NimbyStatus)

func (am *AutoMode) testRunningProcess(nsc chan *NimbyStatus) {
	temp_nimby := NewNimbyStatus()
	temp_nimby.SetReason("testRunningProcess")
	nsc <- temp_nimby
}

func (am *AutoMode) testCPUUSage(nsc chan *NimbyStatus) {
	temp_nimby := NewNimbyStatus()
	temp_nimby.SetReason("testAnother")
	nsc <- temp_nimby
}

func (am *AutoMode) testUsageDelay() {
	log.Printf("[NIMBY] Day mode usage check ...")

	var c chan *NimbyStatus = make(chan *NimbyStatus)

	checks := []testFunc{
		am.testRunningProcess,
		am.testCPUUSage,
	}

	for n := range checks {
		go checks[n](c)
		go FlushByChannel(c)
	}
}

func (am *AutoMode) StartLoop() {
	log.Printf("[NIMBY] start loop ...")
	log.Printf("[NIMBY]")

	// init if ticker object is null
	if am.ticker == nil {
		//interval := viper.GetInt("nimby.autoMode.usageCheckInterval")
		am.ticker = time.NewTicker(time.Second)
	}

	go func() {
		for {
			select {
			case <-am.tickerChannel:
				log.Printf("[NIMBY] ssssssss loop ...")
				return

			case <-am.ticker.C:
				am.testUsageDelay()
				log.Printf("[NIMBY] aaaaaas..")
				log.Printf("[NIMBY]")
			}
		}
	}()
}

func (am *AutoMode) StopLoop() {
	am.ticker.Stop()
	am.tickerChannel <- true
}

func (am *AutoMode) IsRunning() bool {
	v := <-am.tickerChannel
	return !v
}
