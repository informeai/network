package network

import (
	"log"
	"testing"
)

//go test -v -run ^TestRun
func TestRun(t *testing.T) {
	stdout, stderr, err := run("echo", "Hello Informeai!")
	if err != nil {
		t.Errorf("TestRun(): got -> %v, want: nil", err)
	}
	log.Println(stdout)
	log.Println(stderr)
}
