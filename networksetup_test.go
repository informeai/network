package network

import (
	"log"
	"testing"
)

//go test -v -run ^TestNewNetWorkSetup
func TestNewNetWorkSetup(t *testing.T) {
	n := NewNetWorkSetup()
	if n == nil {
		t.Error("TestNewNetWorkSetup(): got -> nil, want: NetWorkSetup{}")
	}
	log.Println(n)
}

//go test -v -run ^TestListNetWorkServiceOrder
func TestListNetWorkServiceOrder(t *testing.T) {
	n := NewNetWorkSetup()
	stdout, err := n.ListNetWorkServiceOrder()
	if err != nil {
		t.Errorf("TestListNetWorkServiceOrder(): got -> %v, want: nil", err)
	}
	log.Println(stdout)
}

//go test -v -run ^TestListAllNetWorkServices
func TestListAllNetWorkServices(t *testing.T) {
	n := NewNetWorkSetup()
	stdout, err := n.ListAllNetWorkServices()
	if err != nil {
		t.Errorf("TestListAllNetWorkServices(): got -> %v, want: nil", err)
	}
	log.Println(stdout)
}

//go test -v -run ^TestListAllHardwarePorts
func TestListAllHardwarePorts(t *testing.T) {
	n := NewNetWorkSetup()
	stdout, err := n.ListAllHardwarePorts()
	if err != nil {
		t.Errorf("TestListAllHardwarePorts(): got -> %v, want: nil", err)
	}
	log.Println(stdout)
}
