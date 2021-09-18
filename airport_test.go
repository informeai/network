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
	resp, err := a.getScan()
	if err != nil {
		t.Errorf("TestGetScan(): got -> %v, want: nil", err)
	}
	log.Println(resp)
}

//go test -v -run ^TestGetInfo
func TestGetInfo(t *testing.T) {
	a := NewAirport()
	resp, err := a.getInfo()
	if err != nil {
		t.Errorf("TestGetInfo(): got -> %v, want: nil", err)
	}
	log.Println(resp)
}

//go test -v -run ^TestParseWifi
func TestParseWifi(t *testing.T) {
	a := NewAirport()
	wifis, err := a.parseWifi()
	if err != nil {
		t.Errorf("TestParseWifi(): got -> %v, want: nil", err)
	}
	log.Println(wifis)
}
