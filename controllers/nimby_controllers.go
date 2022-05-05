package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/OlivierArgentieri/go_killprocess/models"
	"github.com/OlivierArgentieri/go_killprocess/responses"
	"github.com/spf13/viper"
)

func (server *Server) SetNimbyStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received SetNimbyStatus request\n")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	temp_store := models.GetStoreInstance()

	currentStatus := temp_store.NimbyStatus
	receiveStatus := models.NewNimbyStatus()
	err = json.Unmarshal(body, &receiveStatus)
	if err != nil {
		log.Printf("[NIMBY] Setting Nimby value \n")
	}

	log.Printf("Update nimby status from: %t, %s, %s", currentStatus.GetValue(), currentStatus.GetMode(), currentStatus.GetReason())
	currentStatus.Merge(receiveStatus)
	log.Printf("to: %t, %s, %s", currentStatus.GetValue(), currentStatus.GetMode(), currentStatus.GetReason())

	// request local blade
	var real_value = "0"

	// convert true to 1 and false to 0
	if *currentStatus.Value {
		real_value = "1"
	}

	res, err := http.Get(fmt.Sprintf("%s/blade/ctrl?nimby=%s", viper.GetString("nimby.bladeURL"), real_value))

	if err != nil {
		log.Printf("[NIMBY] error setting currentStatus on local blade process \n")
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// read body response of blade status request
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	// response is in json so we need to decode it and convet tu byte
	raw_json := []byte(string(body))

	var dat map[string]interface{}
	// unmarshal
	if err := json.Unmarshal(raw_json, &dat); err != nil {
		panic(err)
	}

	// return response of blade to response of request
	responses.JSON(w, http.StatusOK, dat)
}

func (server *Server) GetBladeStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received kill request\n")
	responses.JSON(w, http.StatusOK, nil)
}
