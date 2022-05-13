package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/OlivierArgentieri/go_killprocess/models"
	"github.com/OlivierArgentieri/go_killprocess/responses"
)

func (server *Server) SetNimbyStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received SetNimbyStatus request\n")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	temp_store := models.GetStoreInstance()
	currentStatus := temp_store.NimbyStatus
	receiveStatus := models.NewNimbyStatus()

	err = json.Unmarshal(body, &receiveStatus)
	if err != nil {
		log.Printf("[NIMBY] Setting Nimby value \n")
		return
	}
	log.Printf("Update nimby status from: %t, %s, %s", currentStatus.GetValue(), currentStatus.GetMode(), currentStatus.GetReason())

	// Test recevied Mode value
	if receiveStatus.GetMode() != models.NIMBY_AUTO && receiveStatus.GetMode() != models.NIMBY_MANUAL {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("invalid state mode, you can only choose manual or auto as mode value"))
		return
	}

	log.Printf("Update nimby status from: %t, %s, %s", currentStatus.GetValue(), currentStatus.GetMode(), currentStatus.GetReason())
	currentStatus.Merge(receiveStatus)
	log.Printf("to: %t, %s, %s", currentStatus.GetValue(), currentStatus.GetMode(), currentStatus.GetReason())

	// flush data to local blade process
	body, err = currentStatus.FlushToNimbyProcess()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// response is in json so we need to decode it and convet tu byte
	raw_json := []byte(string(body))

	var dat map[string]interface{}
	// unmarshal
	if err := json.Unmarshal(raw_json, &dat); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// return response of blade to response of request
	responses.JSON(w, http.StatusOK, dat)
}

func (server *Server) GetBladeStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received kill request\n")
	responses.JSON(w, http.StatusOK, nil)
}

func (server *Server) TestTemp(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received test request\n")

}
