package network

import (
	"bytes"
	"os/exec"
	"strings"
	"time"
)

//Airport is utility for manager wifi-points.
type Airport struct {
	stdout bytes.Buffer
	stderr bytes.Buffer
}

//Wifi is struct with info about access points.
type Wifi struct {
	SSID     string
	BSSID    string
	RSSI     string
	CHANNEL  string
	SECURITY string
}

//NewAirport return new instance of airport
func NewAirport() *Airport {
	return &Airport{}
}

//run execute commands.
func (a *Airport) run(command, param string) (stdout, stderr string, err error) {
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

//getScan return list of networks wifi.
func (a *Airport) getScan() (string, error) {
	stdout, _, err := a.run(cmdAirport, "-s")
	if err != nil {
		return "", err
	}
	return stdout, nil

}

//getInfo return current wireless status, e.g. signal info, BSSID, port type etc.
func (a *Airport) getInfo() (string, error) {
	stdout, _, err := a.run(cmdAirport, "-I")
	if err != nil {
		return "", err
	}
	return stdout, nil
}

//parseWifi return slice of Wifi parsed.
func (a *Airport) parseWifi() ([]Wifi, error) {
	var wifis []Wifi
	stdout, err := a.getScan()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(stdout, "\n")
	lines = lines[1:]
	for _, v := range lines {
		line := strings.TrimSpace(v)
		l := strings.Split(line, " ")
		if len(l) > 6 {
			var sec string
			if strings.Contains(l[4], ",") {
				sec = l[11]
			} else {
				sec = l[13]
			}
			w := Wifi{
				SSID:     l[0],
				BSSID:    l[1],
				RSSI:     l[2],
				CHANNEL:  l[4],
				SECURITY: sec,
			}
			wifis = append(wifis, w)

		}

	}
	return wifis, nil
}
