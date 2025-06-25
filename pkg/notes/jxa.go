package notes

import (
	"fmt"
	"os"
	"os/exec"
)

func executeJXA(script string) ([]byte, error) {
	tempFile, err := os.CreateTemp("", "jxa-*.jxa")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if _, err := tempFile.WriteString(script); err != nil {
		return nil, fmt.Errorf("failed to write script to temp file: %w", err)
	}

	if err := tempFile.Close(); err != nil {
		return nil, fmt.Errorf("failed to close temp file: %w", err)
	}

	cmd := exec.Command("osascript", "-l", "JavaScript", tempFile.Name())
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("JXA script failed: %s", string(exitError.Stderr))
		}
		return nil, fmt.Errorf("failed to execute JXA script: %w", err)
	}
	return output, nil
}

