package suspend

import (
	"log"
	"os/exec"
)

type keyType = int32

const (
	h keyType = 0
	s keyType = 1
)

// See https://docs.microsoft.com/en-us/windows/win32/api/powrprof/nf-powrprof-setsuspendstate
var suspendParameters = map[keyType]string{
	h: "1,1,0",
	s: "0,1,0",
}

func rundllHibernate() {
	run(h)
}

func run(suspendParam keyType) {
	cmd := exec.Command("rundll32.exe", "powrprof.dll,SetSuspendState", suspendParameters[suspendParam])
	err := cmd.Start()

	if err != nil {
		log.Fatalf("An error occured during hibernation: %s\n", err)
	}
}
