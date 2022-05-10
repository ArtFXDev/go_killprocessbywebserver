package models

import (
	"log"

	"github.com/OlivierArgentieri/go_killprocess/utils"
)

type AutoMode struct {
	// private
	si *utils.SetInterval

	// public
	Name     string
	Interval string
}
type testFunc func(chan *NimbyStatus)

func (am *AutoMode) testRunningProcess(nsc chan *NimbyStatus) {
	log.Println("testRunningProcess")
	temp_nimby := NewNimbyStatus()
	temp_nimby.SetReason("testRunningProcess")
	nsc <- temp_nimby
}

func (am *AutoMode) testCPUUSage(nsc chan *NimbyStatus) {
	log.Println("testCPUUSage")

	temp_nimby := NewNimbyStatus()
	temp_nimby.SetReason("testAnother")
	nsc <- temp_nimby
}

func (am *AutoMode) testUsageDelay() {
	log.Printf("testUsageDelay")

	var c chan *NimbyStatus = make(chan *NimbyStatus)

	checks := []testFunc{
		am.testRunningProcess,
		am.testCPUUSage,
	}

	for n := range checks {
		go checks[n](c)
		FlushByChannel(c)
	}
}

func (am *AutoMode) IsValid() bool {
	return am.si != nil
}

func (am *AutoMode) StartLoop() {

	if am.IsValid() && am.si.IsRunning() {
		log.Println("already running")
		return
	}

	am.si = &utils.SetInterval{}
	am.si.Start(am.testUsageDelay, 1000, false)
}

func (am *AutoMode) StopLoop() {

	if !am.IsValid() {
		return
	}

	// already stopped
	if !am.si.IsRunning() {
		return
	}

	log.Println("Stopping loop")

	am.si.Stop()
}
