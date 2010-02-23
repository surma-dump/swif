package main

import (
	"os"
	"swif"
)


func main() {
	f, err := os.Open("swif.conf.xml", os.O_RDONLY, 0)
	if err != nil {
		panic("Could not open config file\n", err.String())
	}

	s := swif.NewSwif()
	err = s.ReadConfig(f)
	if err != nil {
		panic("Reading configuration failed\n", err.String())
	}
	err = s.Start()
	if err != nil {
		panic("Could not start daemon\n", err.String())
	}
}
