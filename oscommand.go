package dangerous

import (
	"os/exec"
)

func RunCommand(userInput string) (string, error) {
	// 🚨 SECURITY RISK: using untrusted input directly in command args
	cmd := exec.Command("sh", "-c", "echo "+userInput)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func SafeRunCommand(userInput string) (string, error) {
	// ✅ SAFE: properly pass user input as an argument
	cmd := exec.Command("echo", userInput)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
