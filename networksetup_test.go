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
