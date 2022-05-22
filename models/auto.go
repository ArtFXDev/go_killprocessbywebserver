package models

import (
	"log"

	"github.com/OlivierArgentieri/go_killprocess/utils"
	"github.com/spf13/viper"
)

type AutoMode struct {
	// private
	si *utils.SetInterval
	c  chan *NimbyStatus

	// public
	Name     string
	Interval string
}
type testFunc func(chan *NimbyStatus)

func (am *AutoMode) initAutoMode() {
	am.si = &utils.SetInterval{}
	am.c = make(chan *NimbyStatus)
}

func (am *AutoMode) testRunningProcess(nsc chan *NimbyStatus) {
	log.Println("testRunningProcess")

	result, err := utils.GetRunningProcess()
	if err != nil {
		log.Printf("Error test usage: %s", err)
		return
	}
	log.Print(result)
	temp_nimby := NewNimbyStatus()
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
		log.Printf("High CPU USAGE: turn off nimby")
		temp_nimby.SetReason("High CPU Usage")
		temp_nimby.SetMode(NIMBY_MANUAL)
		temp_nimby.SetValue(false)
	} else {
		log.Printf("LOW CPU USAGE: turn on nimby")
		temp_nimby.SetReason("Low CPU Usage")
		temp_nimby.SetMode(NIMBY_AUTO)
		temp_nimby.SetValue(true)
	}

	nsc <- temp_nimby
}

func (am *AutoMode) testUsageDelay() {
	log.Printf("testUsageDelay")
	checks := []testFunc{
		am.testCPUUSage,
		am.testRunningProcess,
	}

	for n := range checks {
		go checks[n](am.c)
		FlushByChannel(am.c)
	}
}

func (am *AutoMode) IsValid() bool {
	return am.si != nil
}

func (am *AutoMode) StartLoop() {

	if !am.IsValid() {
		am.initAutoMode()
	}

	if am.si.IsRunning() {
		log.Println("already running")
		return
	}

	delay := viper.GetInt("nimby.automode.usageCheckInterval")
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
