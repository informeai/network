package network

import (
	"log"
	"testing"
)

//go test -v -run ^TestNewAirport
func TestNewAirport(t *testing.T) {
	a := newAirport()
	if a == nil {
		t.Error("TestNewAirport(): got -> nil, want: airport{}")
	}
	log.Println(a)

}

//go test -v -run ^TestRun
func TestRun(t *testing.T) {
	a := newAirport()
	stdout, stderr, err := a.run(cmdAirport, "-I")
	if err != nil {
		t.Errorf("TestRun(): got -> %v, want: nil", err)
	}
	log.Println(stdout)
	log.Println(stderr)
}

//go test -v -run ^TestScan
func TestScan(t *testing.T) {
	a := newAirport()
	resp, err := a.scan()
	if err != nil {
		t.Errorf("TestScan(): got -> %v, want: nil", err)
	}
	log.Println(resp)
}

//go test -v -run ^TestGetInfo
func TestGetInfo(t *testing.T) {
	a := newAirport()
	resp, err := a.getInfo()
	if err != nil {
		t.Errorf("TestGetInfo(): got -> %v, want: nil", err)
	}
	log.Println(resp)
}
