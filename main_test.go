package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMainWithNoArgs(t *testing.T) {
	cmd := exec.Command("go", "run", ".")
	output, err := cmd.CombinedOutput()

	if err == nil {
		t.Fatal("エラーが起こるはずなのでおかしいです")
	}

	if !strings.Contains(string(output), "Usage") {
		t.Fatal("Usage を表示してあげてください:", string(output))
	}
}
