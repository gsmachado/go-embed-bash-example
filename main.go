package main

import (
	_ "embed"
	"fmt"
	"os/exec"
	"strings"
)

//go:embed scripts/script.sh
var script []byte

func main() {

	// check for `bash` in the system
	_, err := exec.LookPath("bash")
	if err != nil {
		fmt.Printf("bash is not available: %v\n", err)
		return
	}

	// -s: read commands from the standard input
	// -: used as a placeholder for the end of options
	c := exec.Command("bash", "-s", "-", "test-param1", "test-param2")
	c.Stdin = strings.NewReader(string(script))

	b, e := c.Output()
	if e != nil {
		exitError, ok := e.(*exec.ExitError)
		if ok {
			fmt.Printf("Script exited with non-zero status code: %d\n", exitError.ExitCode())
		} else {
			// If the error is of a different type, it means that
			// there was an error in running the process itself,
			// not that the process ran and exited with an error.
			// This could be due to reasons like the command not
			// being found, the process being killed due to a signal, etc.
			fmt.Printf("Failed to run script: %v\n", e)
		}
		return
	}
	fmt.Println(string(b))
}
