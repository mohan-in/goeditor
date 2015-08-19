package main

import (
	"bytes"
	"os/exec"
	"strings"
)

func autoComplete(file []byte, offset string) []string {
	cmd := exec.Command("C:\\wrk\\bin\\gocode.exe", "autocomplete", "c"+offset)

	cmd.Stdin = bytes.NewReader(file)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		logger.Fatal(err)
	}

	return strings.Split(string(out.Bytes()), "\n")
}
