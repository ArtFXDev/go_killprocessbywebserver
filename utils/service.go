package utils

import (
	"os/exec"
)

func RestartService(serviceName string) error {
	restartService := exec.Command("Get-Service", "-Name", "net*", "|", "Where-Object", "{$_.Status -eq", serviceName, "}", "|", "Restart-Service")
	err := restartService.Run()
	if err != nil {
		return err
	}
	return nil
}
