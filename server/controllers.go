package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/OlivierArgentieri/go_killprocess/responses"
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

	log.Println("Try to kill pid:", strconv.Itoa(int(pid)))
	kill := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(int(pid)))
	err = kill.Run()

	if err != nil {
		responses.JSON(w, http.StatusOK, "Error when trying to kill process, pls verify the requested PID")
		log.Printf("Error when trying to kill process, pls verify the requested PID")
		return
	}

	log.Printf("Kill successfull")
	responses.JSON(w, http.StatusOK, "Success")
}

func (server *Server) GetProcesses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received get processes request\n")

	const cmd_powershell = `get-wmiObject Win32_PerfFormattedData_PerfProc_Process | ? {$_.Name -notlike '*_Total'} | ? {$_.Name -notlike '*Idle'} | ? {$_.PercentProcessorTime -gt '0'} | sort-object PercentProcessorTime -desc | select @{N='Name';E={$_.Name}}, @{N='CPU';E={$_.PercentProcessorTime}}, @{N='RAM';E={([math]::Round($_.WorkingSetPrivate/1Mb,2))}}, @{N='PID';E={$_.CreatingProcessID}} | ConvertTo-Json`

	get_processes := exec.Command("powershell.exe", cmd_powershell)
	log.Print(get_processes)

	stdout, stderr := get_processes.CombinedOutput()

	//err := get_processes.Run()

	if stderr != nil {
		responses.JSON(w, http.StatusInternalServerError, "Error when trying to get list of process.")
		log.Printf("Error when trying to get list of process.")
		log.Printf("err processes: %v ", stderr)
		log.Printf("log processes: %v", string(stdout))
		return
	}

	type ProcessRow struct {
		Name string
		CPU  float32
		RAM  float32
		PID  float32
	}
	var rows []ProcessRow
	json.Unmarshal([]byte(stdout), &rows)
	if stderr != nil {
		log.Printf("Error when Unmarshal json: %v", stderr)
		responses.JSON(w, http.StatusInternalServerError, "Error when Unmarshal json: ")
	}
	log.Printf("Ok get processes: %v", rows)
	responses.JSON(w, http.StatusOK, rows)
}
