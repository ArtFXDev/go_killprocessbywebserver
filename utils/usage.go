package utils

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetCPUUsage() (int, error) {
	result, err := exec.Command("wmic", "cpu", "get", "loadpercentage").Output()
	if err != nil {
		return 0, err
	}
	beforeint := strings.ReplaceAll(string(result), "loadpercentage", "")

	intValue, err := strconv.Atoi(beforeint)
	if err != nil {
		return 0, err
	}

	return intValue, nil

}
