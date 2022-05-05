package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// type NimbyStatus interface {}

type NimbyStatus struct {
	Value  *bool   `json:"value"`
	Mode   *string `json:"mode"`
	Reason *string `json:"reason"`
}

func NewNimbyStatus() *NimbyStatus {
	return &NimbyStatus{&[]bool{true}[0], &[]string{"auto"}[0], &[]string{"Default status"}[0]}
}

func (status *NimbyStatus) Merge(otherStatus *NimbyStatus) {
	if otherStatus.Value != nil {
		status.Value = otherStatus.Value
	}

	if otherStatus.Mode != nil {
		status.Mode = otherStatus.Mode
	}

	if otherStatus.Reason != nil {
		status.Reason = otherStatus.Reason
	}
}

func (status *NimbyStatus) GetValue() bool {
	return *status.Value
}

func (status *NimbyStatus) GetMode() string {
	return *status.Mode
}

func (status *NimbyStatus) GetReason() string {
	return *status.Reason
}

func (status *NimbyStatus) SetValue(v bool) {
	status.Value = &[]bool{v}[0]
	GetStoreInstance().NimbyStatus.Value = &[]bool{v}[0]
}

func (status *NimbyStatus) SetMode(v string) {
	// todo make mode as Enum
	status.Mode = &[]string{v}[0]
	GetStoreInstance().NimbyStatus.Mode = &[]string{v}[0]
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
