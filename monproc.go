package main

/*
	Stack Overflow - How to get total CPU usage from /proc/pid/stat?

	http://man7.org/linux/man-pages/man5/proc.5.html
		-> search 'uptime' and 'stat'
*/

// add wtfBUFF in a struct to share output of

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Monproc -  a process monitor for Debian Linux Distros.
type Monproc interface {
	// calcCPU
	// setState()
	getUptime()
	rFile(p string) []byte
}

type process struct {
	Monproc
	path      string
	uptime    int
	utime     uint16
	stime     uint16
	cutime    uint16
	cstime    uint16
	starttime int
	hertz     int
	state     string
	pid       int
}

func (mp *process) setState(s rune) {
	statemap := map[rune]string{
		'R': "Running",
		'S': "Sleeping in an interruptible wait",
		'D': "Waiting in uninterruptible disk sleep",
		'Z': "Zombie",
		'T': "Stopped (on a signal)",
		't': "Tracing stop",
		'X': "Dead",
		'x': "Dead",
		'K': "Wakekill",
		'W': "Waking",
		'P': "Parked",
	}
	mp.state = statemap[s]
}

func (mp process) rFile(p string) []byte {
	content, _ := ioutil.ReadFile(mp.path + p)
	return content
}

func (mp *process) getUptime() {
	uptimeOUT := mp.rFile("uptime")
	fmt.Println(uptimeOUT)
}

// GetProcesses - get percentage of CPU usage per running process
func GetProcesses() {
	var path string = "/proc/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Read error. Are you root?")
		os.Exit(1)
	}
	for i, pid := range files {
		intPID, err := strconv.Atoi(pid.Name())
		if err != nil {
			continue
		}

		// ** remove this ** //
		if i < 3000 {
			continue
		}
		// ***************** //

		var monproc Monproc
		monproc = &process{path: path + pid.Name(), pid: intPID}
		monproc.getUptime()

		// ** remove this when goroutines added ** //
		break
		// *************************************** //
	}
}

func main() {
	GetProcesses()
}
