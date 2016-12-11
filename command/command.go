package command

import (
	"bytes"
	"log"
	"os/exec"
	"syscall"
)

type Result struct {
	StdOut   string
	StdErr   string
	ExitCode int
}

const defaultFailedCode = 1

// exit code checking in go:
// 1. try to get the exit code
// 2. This will happen (in OSX) if `name` is not available in $PATH,
// in this situation, exit code could not be get, and stderr will be
// empty string very likely, so we use the default fail code, and format err
// to string and set to stderr
// 3. success, exitCode should be 0 if go is ok
func Run(name string, args ...string) (stdout string, stderr string, exitCode int) {
	log.Println("run command:", name, args)
	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}

	return stdout, stderr, exitCode
}
