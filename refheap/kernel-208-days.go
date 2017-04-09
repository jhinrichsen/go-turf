package main

import (
	"fmt"
	"log"
	"syscall"
	"strconv"
	"time"
)

// Calculate when your Linux kernel is vulnerable to crashes
// Used for critical linux kernel bug 'Time Stamp Counter'
// Novell: SuSE document no 7009834
func main() {
	var si syscall.Sysinfo_t
	err := syscall.Sysinfo(&si)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	// Last boot time is now - uptime [s]
	lastBoot := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second() - int(si.Uptime), now.Nanosecond(), now.Location())

	// Uptime is seconds since last boot
	fmt.Printf("Now: %s, last boot: %s, uptime: %s\n", now, lastBoot, now.Sub(lastBoot))

	// Critical uptime begins at 208 days
	criticalUptime, err := time.ParseDuration(strconv.Itoa(208*24) + "h")
	if err != nil {
		log.Fatal(err)
	}
	vulnerable := lastBoot.Add(criticalUptime)
	fmt.Printf("System starts to be vulnerable at %s\n", vulnerable)
}

