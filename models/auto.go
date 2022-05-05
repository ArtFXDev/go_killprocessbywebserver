package models

import "log"

type AutoMode struct {
	Name     string
	Interval string
}
type testFunc func(chan *NimbyStatus)

func (am *AutoMode) testRunningProcess(ns chan *NimbyStatus) {
	temp_nimby := NewNimbyStatus()
	ns <- temp_nimby
}

func (am *AutoMode) testAnother(ns chan *NimbyStatus) {
	temp_nimby := NewNimbyStatus()
	ns <- temp_nimby
}

func (am *AutoMode) testUsageDelay() {
	log.Printf("[NIMBY] Day mode usage check ...")

	// var c chan *NimbyStatus = make(chan *NimbyStatus)

	checks := []testFunc{
		am.testRunningProcess,
		am.testAnother,
	}

	for n := range checks {
		go checks[n](nil)
	}
}
