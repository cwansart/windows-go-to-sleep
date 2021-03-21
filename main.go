package main

import (
	"log"
	"os/exec"
)

// See https://docs.microsoft.com/en-us/windows/win32/api/powrprof/nf-powrprof-setsuspendstate
var suspendParameters = map[string]string{
	"hibernate": "1,1,0",
	"sleep":     "0,1,0",
}

func main() {
	startHibernation()
}

func startHibernation() {
	cmd := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", suspendParameters["hibernate"])
	err := cmd.Start()

	if err != nil {
		log.Fatalf("An error occured during hibernation: %s\n", err)
	}
}
