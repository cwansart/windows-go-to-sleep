package main

import (
	"flag"
	"log"
	"time"

	"github.com/cwansart/windows-go-to-sleep/suspend"
)

var (
	flagUnit  = flag.String("unit", "minutes", "unit which will be in combination with the value param. Possible values: seconds, minutes, hours")
	flagValue = flag.Uint("value", 60, "value in which the system will go into hibernation")
)

func main() {
	flag.Parse()
	log.Printf("System will go into hibernation in %d %s\n", *flagValue, *flagUnit)

	durationType := determineDurationType(*flagUnit)
	timer := time.NewTimer(time.Duration(*flagValue) * durationType)
	<-timer.C

	suspend.Hibernate()
}

func determineDurationType(unit string) time.Duration {
	var t time.Duration

	switch unit {
	case "seconds":
		t = time.Second
	case "minutes":
		t = time.Minute
	case "hours":
		t = time.Hour
	default:
		log.Fatalf("Invalid value for parameter 'unit': %s\n", unit)
	}

	return t
}
