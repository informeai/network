package network

import (
	"bytes"
	"os/exec"
	"time"
)

//var buffers.
var stdout, stderr bytes.Buffer

//functions for execution of utilitys.

//run execute commands.
func run(command, param string) (string, string, error) {
	var cmdOut, cmdErr string
	cmd := exec.Command(command, param)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		return "", "", err
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeDuration):
		err = cmd.Process.Kill()
	case err = <-done:
		cmdOut = stdout.String()
		cmdErr = stderr.String()
	}
	return cmdOut, cmdErr, nil
}
