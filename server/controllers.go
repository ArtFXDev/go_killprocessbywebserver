package server

import (
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
