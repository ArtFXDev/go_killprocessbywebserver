package utils

import (
	"os/exec"
)

func RestartService(serviceName string) error {
	restartService := exec.Command("powershell.exe", "Restart-Service", "-Name", serviceName)
	err := restartService.Run()
	if err != nil {
		return err
	}
	return nil
}
