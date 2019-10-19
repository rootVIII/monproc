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
	"strings"
)

// Monproc -  a process monitor for Debian Linux Distros.
type Monproc interface {
	// calcCPU
	// setState()
	getUptime()
	rFile(p string) string
}

type process struct {
	Monproc
	path      string
	uptime    float64
	utime     int
	stime     int
	cutime    int
	cstime    int
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

func (mp process) rFile(p string) string {
	content, _ := ioutil.ReadFile(mp.path + p)
	return string(content)
}

func (mp *process) getUptime() {
	uptimeOUT := mp.rFile("uptime")
	mp.uptime, _ = strconv.ParseFloat(strings.Split(uptimeOUT, " ")[0], 64)
	fmt.Println(mp.uptime)
}

// GetProcesses - get percentage of CPU usage per running process
func GetProcesses() {
	var path string = "/proc/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Read error. Are you root?")
		os.Exit(1)
	}
	for _, pid := range files {
		PID, err := strconv.Atoi(pid.Name())
		if err != nil {
			continue
		}

		// ** remove this ** //
		if PID < 2000 {
			continue
		}
		// ***************** //
		var monproc Monproc
		monproc = &process{path: path, pid: PID}
		monproc.getUptime()

		// ** remove this when goroutines added ** //
		break
		// *************************************** //
	}
}

func main() {
	GetProcesses()
}
