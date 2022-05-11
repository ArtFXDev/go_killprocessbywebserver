package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// type NimbyStatus interface {}
type NimbyStatusMode string

const (
	NIMBY_AUTO   NimbyStatusMode = "auto"
	NIMBY_MANUAL NimbyStatusMode = "manual"
)

type NimbyStatus struct {
	Value  *bool            `json:"value"`
	Mode   *NimbyStatusMode `json:"mode"`
	Reason *string          `json:"reason"`

	AutoMode *AutoMode
}

func NewNimbyStatus() *NimbyStatus {
	instance := &NimbyStatus{
		&[]bool{true}[0],
		&[]NimbyStatusMode{NIMBY_AUTO}[0],
		&[]string{"Default status"}[0],
		&AutoMode{},
	}
	//	instance.AutoMode.StartLoop()
	return instance
}

func (status *NimbyStatus) Merge(otherStatus *NimbyStatus) {
	log.Println("[NIMBY] call Merge")

	if otherStatus.Value != nil {
		status.SetValue(*otherStatus.Value)
	}

	if otherStatus.Mode != nil {
		status.SetMode(*otherStatus.Mode)
	}

	if otherStatus.Reason != nil {
		status.SetReason(*otherStatus.Reason)
	}
}

func (status *NimbyStatus) GetValue() bool {
	return *status.Value
}

func (status *NimbyStatus) GetMode() NimbyStatusMode {
	return *status.Mode
}

func (status *NimbyStatus) GetReason() string {
	return *status.Reason
}

func (status *NimbyStatus) SetValue(v bool) {
	status.Value = &[]bool{v}[0]
	GetStoreInstance().NimbyStatus.Value = &[]bool{v}[0]
}

func (status *NimbyStatus) SetMode(v NimbyStatusMode) {
	log.Println("[NIMBY] call setMode")

	status.Mode = &[]NimbyStatusMode{v}[0]
	GetStoreInstance().NimbyStatus.Mode = &[]NimbyStatusMode{v}[0]

	if *(status.Mode) == NIMBY_AUTO {
		status.AutoMode.StartLoop()
	} else {
		status.AutoMode.StopLoop()
	}
}

func (status *NimbyStatus) SetReason(v string) {
	status.Reason = &[]string{v}[0]
	GetStoreInstance().NimbyStatus.Reason = &[]string{v}[0]
}

func (status *NimbyStatus) FlushToNimbyProcess() ([]byte, error) {

	// request local blade
	var real_value = "0"

	// convert true to 1 and false to 0
	if *status.Value {
		real_value = "1"
	}

	res, err := http.Get(fmt.Sprintf("%s/blade/ctrl?nimby=%s", viper.GetString("nimby.bladeURL"), real_value))

	if err != nil {
		log.Println("[NIMBY] error setting currentStatus on local blade process")
		return []byte{}, err
	}

	// read body response of blade status request
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("[NIMBY] error setting currentStatus on local blade process")
		return []byte{}, err
	}

	return body, nil
}

func FlushByChannel(nsc chan *NimbyStatus) {
	receive_NimbyStatus := <-nsc
	log.Println("[NIMBY] FlushByChannel receive from goroutine to local nimby process")

	receive_NimbyStatus.FlushToNimbyProcess()
	temp_store := GetStoreInstance()
	temp_store.NimbyStatus.Merge(receive_NimbyStatus)

	log.Printf("Update nimby status from: %t, %s, %s", temp_store.NimbyStatus.GetValue(), temp_store.NimbyStatus.GetMode(), temp_store.NimbyStatus.GetReason())
	log.Printf("To: %t, %s, %s", receive_NimbyStatus.GetValue(), receive_NimbyStatus.GetMode(), receive_NimbyStatus.GetReason())

}
