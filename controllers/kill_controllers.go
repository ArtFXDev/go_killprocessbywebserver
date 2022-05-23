package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/OlivierArgentieri/go_killprocess/responses"
	"github.com/OlivierArgentieri/go_killprocess/utils"
	"github.com/gorilla/mux"
)

func (server *Server) KillProcess(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received kill request\n")
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["pid"], 10, 64)
	if err != nil {
		log.Printf("ERROR when trying to parse pid parameter")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	var ipid = int(pid)

	err = utils.Terminate(ipid)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, "Error when trying to kill process, pls verify the requested PID")
		log.Printf("Error when trying to kill process, pls verify the requested PID, error: %v", err)
		return
	}

	log.Printf("Kill successfull")
	responses.JSON(w, http.StatusOK, "Success")
}

func (server *Server) GetProcesses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received get processes request\n")

	var cmd_powershell = ` ps | foreach-Object {@{"Name"=$_.name;"PID"=$_.ID; "RAM"=$_.WS}} | ConvertTo-json`

	get_processes := exec.Command("powershell.exe", cmd_powershell)
	log.Print(get_processes)

	stdout, stderr := get_processes.CombinedOutput()

	if stderr != nil {
		responses.JSON(w, http.StatusInternalServerError, "Error when trying to get list of process.")
		log.Printf("Error when trying to get list of process.")
		log.Printf("err processes: %v ", stderr)
		log.Printf("log processes: %v", string(stdout))
		return
	}

	type ProcessRow struct {
		Name string
		RAM  float32
		PID  float32
	}

	var rows []ProcessRow
	json.Unmarshal([]byte(stdout), &rows)
	if stderr != nil {
		log.Printf("Error when Unmarshal json: %v", stderr)
		responses.JSON(w, http.StatusInternalServerError, "Error when Unmarshal json: ")
		return
	}
	log.Printf("Ok get processes: %v", rows)
	responses.JSON(w, http.StatusOK, rows)
}

func (server *Server) RestartServices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	process_name := vars["name"]
	log.Printf("Try to restart process: %v", process_name)
	err := utils.RestartService(process_name)

	if err != nil {
		log.Printf("Try to restart process: %v", process_name)
		responses.JSON(w, http.StatusInternalServerError, "Error when Unmarshal json: ")
	}
}
