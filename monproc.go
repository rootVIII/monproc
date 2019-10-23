package main

/*
	rootVIII
	monproc - Displays CPU usage

	Intended for Debian Linux Distros
*/

// #include <unistd.h>
// static int cpuSeconds() {
//     return sysconf(_SC_CLK_TCK);
// }
import "C"
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Monproc -  a process monitor for Debian Linux Distros.
type Monproc interface {
	calcCPU()
	getCPUSeconds()
	setState(s rune)
	getUptime()
	getStat()
	rFile(p string) []byte
}

type process struct {
	Monproc
	percentage float64
	uptime     float64
	utime      int
	stime      int
	cutime     int
	cstime     int
	starttime  int
	hertz      int
	state      string
	pid        string
	name       string
	path       string
}

func (mp *process) setState(s rune) {
	statemap := map[rune]string{
		'I': "Idle",
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

func (mp *process) calcCPU() {
	var total int = mp.utime + mp.stime + mp.cutime + mp.cstime
	var sec float64 = mp.uptime - (float64(mp.starttime) / float64(mp.hertz))
	percentage := 100 * ((float64(total) / float64(mp.hertz)) / sec)
	fmt.Printf("%f\n", percentage)
}

func (mp *process) getCPUSeconds() {
	mp.hertz = int(C.cpuSeconds())
}

func (mp process) rFile(p string) []byte {
	content, _ := ioutil.ReadFile(mp.path + p)
	return content
}

func (mp *process) getUptime() {
	uptimeOut := bytes.Split(mp.rFile("uptime"), []byte(" "))
	var uptime float64
	fmt.Fscanf(bytes.NewReader(uptimeOut[0]), "%f", &uptime)
	mp.uptime = uptime
}

func (mp *process) getStat() {
	statOut := strings.Split(string(mp.rFile(mp.pid+"/stat")), " ")
	mp.name = statOut[1][1 : len(statOut[1])-1]
	mp.setState([]rune(statOut[2])[0])
	mp.utime, _ = strconv.Atoi(statOut[13])
	mp.stime, _ = strconv.Atoi(statOut[14])
	mp.cutime, _ = strconv.Atoi(statOut[15])
	mp.cstime, _ = strconv.Atoi(statOut[16])
	mp.starttime, _ = strconv.Atoi(statOut[21])
}

func monProcWrpr(procPath string, pid string) {
	var monproc Monproc
	monproc = &process{path: procPath, pid: pid}
	monproc.getUptime()
	monproc.getCPUSeconds()
	monproc.getStat()
	monproc.calcCPU()
}

// GetProcesses - get percentage of CPU usage per running process
func GetProcesses() {
	var path string = "/proc/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("Read error")
		os.Exit(1)
	}
	// use PID with rest of returned data
	for _, pid := range files {
		_, err := strconv.Atoi(pid.Name())
		if err != nil {
			continue
		}
		monProcWrpr(path, pid.Name())
	}
}

func main() {
	GetProcesses()
}
