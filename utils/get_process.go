package utils

import (
	"os/exec"
	"strings"
)

func GetRunningProcess() ([]string, error) {
	result, err := exec.Command("powershell.exe", "gps | Sort-Object -unique | Format-Table Name -HideTableHeaders").Output()
	if err != nil {
		return []string{}, err
	}

	as_array := strings.Split(string(result), "\n")

	clean_array := []string{}
	for _, str := range as_array {
		str = strings.ReplaceAll(str, " ", "")
		str = strings.ReplaceAll(str, "\n", "")
		str = strings.ReplaceAll(str, "\r", "")
		if str != "" {
			clean_array = append(clean_array, str)
		}
	}

	// get result as array
	return clean_array, err

}
