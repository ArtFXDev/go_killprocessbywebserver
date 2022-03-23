package utils

import (
	"os/exec"
	"strconv"
)

func Terminate(pid int) error {
	kill := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid))
	err := kill.Run()
	if err != nil {
		return err
	}
	return nil
}
