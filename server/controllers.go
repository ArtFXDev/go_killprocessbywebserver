package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/OlivierArgentieri/go_killprocess/responses"
	"github.com/gorilla/mux"
)

func (server *Server) KillProcess(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["pid"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	var obj = map[string]interface{}{}
	err = json.Unmarshal(body, &obj)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	kill := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(int(pid)))
	err = kill.Run()

	if err != nil {
		responses.JSON(w, http.StatusOK, "Error when trying to kill process, pls verify the requested PID")
		return
	}
	responses.JSON(w, http.StatusOK, "Success")
}
