package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func detokenizeAndMask(token string) (string, error) {
	cmd := exec.Command("cte", "detokenize", "--token", token, "--tokengroup", "default", "--format", "N")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("detokenization failed: %v", err)
	}

	detokenizedData := strings.TrimSpace(out.String())

	if len(detokenizedData) > 2 {
		maskedData := detokenizedData[:len(detokenizedData)-2] + "**"
		return maskedData, nil
	}

	return detokenizedData, nil
}

func main() {
	token := "<TOKEN>"

	maskedData, err := detokenizeAndMask(token)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Masked Data:", maskedData)
}
