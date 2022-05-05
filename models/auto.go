package models

import "log"

type AutoMode struct {
	Name     string
	Interval string
}
type testFunc func(chan *NimbyStatus)

func (am *AutoMode) testRunningProcess(nsc chan *NimbyStatus) {
	temp_nimby := NewNimbyStatus()
	temp_nimby.SetReason("testRunningProcess")
	nsc <- temp_nimby
}

func (am *AutoMode) testAnother(nsc chan *NimbyStatus) {
	temp_nimby := NewNimbyStatus()
	temp_nimby.SetReason("testAnother")
	nsc <- temp_nimby
}

func (am *AutoMode) TestUsageDelay() {
	log.Printf("[NIMBY] Day mode usage check ...")

	var c chan *NimbyStatus = make(chan *NimbyStatus)

	checks := []testFunc{
		am.testRunningProcess,
		am.testAnother,
	}

	for n := range checks {
		go checks[n](c)
		go FlushByChannel(c)
	}
}
