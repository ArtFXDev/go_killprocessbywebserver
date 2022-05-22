package utils

import (
	"os/exec"
	"regexp"
	"strconv"
)

func GetCPUUsage() (int, error) {
	result, err := exec.Command("wmic", "cpu", "get", "loadpercentage").Output()
	if err != nil {
		return 0, err
	}

	// get number in result
	r, _ := regexp.Compile(`\d+`)
	onlyValue := r.Find(result)
	intValue, err := strconv.Atoi(string(onlyValue))
	if err != nil {
		return 0, err
	}

	return intValue, nil
}
