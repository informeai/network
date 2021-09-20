package network

import (
	"log"
	"testing"
)

//go test -v -run ^TestNewAirport
func TestNewAirport(t *testing.T) {
	a := NewAirport()
	if a == nil {
		t.Error("TestNewAirport(): got -> nil, want: airport{}")
	}
	log.Println(a)

}

//go test -v -run ^TestRun
func TestRun(t *testing.T) {
	a := NewAirport()
	stdout, stderr, err := a.run(cmdAirport, "-I")
	if err != nil {
		t.Errorf("TestRun(): got -> %v, want: nil", err)
	}
	log.Println(stdout)
	log.Println(stderr)
}

//go test -v -run ^TestGetScan
func TestGetScan(t *testing.T) {
	a := NewAirport()
	resp, err := a.GetScan()
	if err != nil {
		t.Errorf("TestGetScan(): got -> %v, want: nil", err)
	}
	log.Println(resp)
}

//go test -v -run ^TestGetInfo
func TestGetInfo(t *testing.T) {
	a := NewAirport()
	resp, err := a.GetInfo()
	if err != nil {
		t.Errorf("TestGetInfo(): got -> %v, want: nil", err)
	}
	log.Println(resp)
}

//go test -v -run ^TestParseWifi
func TestParseWifi(t *testing.T) {
	a := NewAirport()
	stdout, _ := a.GetScan()
	wifis, err := a.parseWifi(stdout)
	if err != nil {
		t.Errorf("TestParseWifi(): got -> %v, want: nil", err)
	}
	log.Println(wifis)
}

//go test -v -run ^TestParseInfo
func TestParseInfo(t *testing.T) {
	a := NewAirport()
	stdout, _ := a.GetInfo()
	wireless, err := a.parseInfo(stdout)
	if err != nil {
		t.Errorf("TestParseInfo(): got -> %v, want: nil", err)
	}
	log.Println(wireless)
}

//go test -v -run ^TestScan
func TestScan(t *testing.T) {
	a := NewAirport()
	wifis, err := a.Scan()
	if err != nil {
		t.Errorf("TestScan(): got -> %v, want: nil", err)
	}
	log.Println(wifis)
}

//go test -v -run ^TestInfo
func TestInfo(t *testing.T) {
	a := NewAirport()
	wireless, err := a.Info()
	if err != nil {
		t.Errorf("TestInfo(): got -> %v, want: nil", err)
	}
	log.Println(wireless)

}
