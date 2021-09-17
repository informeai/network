package network

import (
	"bytes"
	"os/exec"
	"time"
)

//airport is utility for manager wifi-points.
type airport struct {
	stdout bytes.Buffer
	stderr bytes.Buffer
}

//NewAirport return new instance of airport
func newAirport() *airport {
	return &airport{}
}

//run execute commands.
func (a *airport) run(command, param string) (stdout, stderr string, err error) {
	cmd := exec.Command(command, param)
	cmd.Stdout = &a.stdout
	cmd.Stderr = &a.stderr
	err = cmd.Start()
	if err != nil {
		return
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(timeDuration):
		err = cmd.Process.Kill()
	case err = <-done:
		stdout = a.stdout.String()
		stderr = a.stderr.String()
	}
	return
}

//scan return list of networks wifi.
func (a *airport) scan() (string, error) {
	stdout, _, err := a.run(cmdAirport, "-s")
	if err != nil {
		return "", err
	}
	return stdout, nil

}

//getInfo return current wireless status, e.g. signal info, BSSID, port type etc.
func (a *airport) getInfo() (string, error) {
	stdout, _, err := a.run(cmdAirport, "-I")
	if err != nil {
		return "", err
	}
	return stdout, nil
}
