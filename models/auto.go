package models

import (
	"log"

	"github.com/OlivierArgentieri/go_killprocess/utils"
	"github.com/spf13/viper"
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
	usage, err := utils.GetCPUUsage()
	if err != nil {
		log.Printf("Error test usage: %s", err)
		return
	}

	if usage > viper.GetInt("nimby.automode.maxCPUUsage") {
		log.Printf("High CPU USAGE: turn off nimby %s", err)
		temp_nimby.SetReason("High CPU Usage")
		temp_nimby.SetMode(NIMBY_MANUAL)
		temp_nimby.SetValue(false)
	} else {
		log.Printf("LOW CPU USAGE: turn on nimby %s", err)
		temp_nimby.SetReason("Low CPU Usage")
		temp_nimby.SetMode(NIMBY_AUTO)
		temp_nimby.SetValue(true)
	}

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

	if !am.IsValid() {
		am.si = &utils.SetInterval{}
	}

	if am.si.IsRunning() {
		log.Println("already running")
		return
	}

	delay := viper.GetInt("nimby.automode.usageCheckInterval")
	log.Println(delay)
	am.si.Start(am.testUsageDelay, delay, true)
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
