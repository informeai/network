package main

import (
	"fmt"
	"github.com/informeai/network"
	"log"
	"os"
)

//usage os params.
var usage = `Network:
Wrapper for network management on MacOS using go language.
USAGE:
network [PARAMS] ARGUMENTS
PARAMS:
	--help -h    help commands.
	--scan -s    scaning by network wireless.
	--info -i    info from current wireless.
`

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(usage)
		os.Exit(0)
	}
	result, err := parseArgs(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)

}

//parseArgs execute parse of args.
func parseArgs(args []string) (string, error) {
	switch args[1] {
	case "--help", "-h":
		fmt.Println(usage)
	case "--scan", "-s":
		var resul string
		a := network.NewAirport()
		resul, err := a.GetScan()
		if err != nil {
			return "", err
		}
		return resul, nil
	case "--info", "-i":
		a := network.NewAirport()
		resul, err := a.GetInfo()
		if err != nil {
			return "", err
		}
		return resul, nil
	}
	return usage, nil
}
