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

//Wireless is struct with information the current connection.
type Wireless struct {
	AgrCtlRSSI      string
	AgrExtRSSI      string
	AgrCtlNoise     string
	AgrExtNoise     string
	State           string
	Mode            string
	LastTxRate      string
	MaxRate         string
	LastAssocStatus string
	Auth802         string
	AuthLink        string
	BSSID           string
	SSID            string
	MCS             string
	Channel         string
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

//GetScan return list of networks wifi.
func (a *Airport) GetScan() (string, error) {
	stdout, _, err := a.run(cmdAirport, "-s")
	if err != nil {
		return "", err
	}
	return stdout, nil

}

//GetInfo return current wireless status, e.g. signal info, BSSID, port type etc.
func (a *Airport) GetInfo() (string, error) {
	stdout, _, err := a.run(cmdAirport, "-I")
	if err != nil {
		return "", err
	}
	return stdout, nil
}

//parseWifi return slice of Wifi parsed.
func (a *Airport) parseWifi(stdout string) ([]Wifi, error) {
	var wifis []Wifi
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

//parseInfo return wireles current formated.
func (a *Airport) parseInfo(stdout string) (Wireless, error) {
	var wireless Wireless
	lines := strings.Split(strings.TrimSpace(stdout), "\n")
	wireless = Wireless{
		AgrCtlRSSI:      lines[0][strings.Index(lines[0], ":")+1:],
		AgrExtRSSI:      lines[1][strings.Index(lines[1], ":")+1:],
		AgrCtlNoise:     lines[2][strings.Index(lines[2], ":")+1:],
		AgrExtNoise:     lines[3][strings.Index(lines[3], ":")+1:],
		State:           lines[4][strings.Index(lines[4], ":")+1:],
		Mode:            lines[5][strings.Index(lines[5], ":")+1:],
		LastTxRate:      lines[6][strings.Index(lines[6], ":")+1:],
		MaxRate:         lines[7][strings.Index(lines[7], ":")+1:],
		LastAssocStatus: lines[8][strings.Index(lines[8], ":")+1:],
		Auth802:         lines[9][strings.Index(lines[9], ":")+1:],
		AuthLink:        lines[10][strings.Index(lines[10], ":")+1:],
		BSSID:           lines[11][strings.Index(lines[11], ":")+1:],
		SSID:            lines[12][strings.Index(lines[12], ":")+1:],
		MCS:             lines[13][strings.Index(lines[13], ":")+1:],
		Channel:         lines[14][strings.Index(lines[14], ":")+1:],
	}
	return wireless, nil
}

//Scan execute scanning e return list of wifi.
func (a *Airport) Scan() ([]Wifi, error) {
	stdout, err := a.GetScan()
	if err != nil {
		return nil, err
	}
	wifis, err := a.parseWifi(stdout)
	if err != nil {
		return nil, err
	}
	return wifis, nil
}

//Info return status, signal, port ... from current wireless.
func (a *Airport) Info() (Wireless, error) {
	stdout, err := a.GetInfo()
	if err != nil {
		return Wireless{}, err
	}
	wireless, err := a.parseInfo(stdout)
	if err != nil {
		return Wireless{}, err
	}
	return wireless, nil
}
