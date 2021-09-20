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
