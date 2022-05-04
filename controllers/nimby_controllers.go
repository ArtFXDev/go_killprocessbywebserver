package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/OlivierArgentieri/go_killprocess/responses"
)

type CurrentStatus struct {
	Value  *bool   `json:"value"`
	Mode   *string `json:"mode"`
	Reason *string `json:"reason"`
}

func (server *Server) SetNimbyStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received SetNimbyStatus request\n")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	currentStatus := CurrentStatus{}
	err = json.Unmarshal(body, &currentStatus)
	if err != nil {
		log.Printf("[NIMBY] Setting Nimby value \n")
	}

	if currentStatus.Value != nil {
		log.Printf("[NIMBY] Setting Nimby value to: %t \n", *currentStatus.Value)
	}

	if currentStatus.Mode != nil {
		log.Printf("[NIMBY] Setting Nimby value to: %s \n", *currentStatus.Mode)
	}

	if currentStatus.Reason != nil {
		log.Printf("[NIMBY] Setting Nimby value to: %s \n", *currentStatus.Reason)
	}

	// request local blade
	var real_value = "0"

	// convert true to 1 and false to 0
	if *currentStatus.Value {
		real_value = "1"
	}

	// route for local blade (todo: read from config file)
	res, err := http.Get(fmt.Sprintf("http://localhost:9005/blade/ctrl?nimby=%s", real_value))

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
