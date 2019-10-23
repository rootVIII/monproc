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
	// calcCPU
	getCPUSeconds()
	setState(s rune)
	getUptime()
	getStat()
	rFile(p string) []byte
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
	pid       string
	name      string
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

	fmt.Printf("%s\n", statOut)
	fmt.Printf("uptime: %f\n", mp.uptime)
	fmt.Println("name: " + mp.name)
	fmt.Println("state: " + mp.state)
	fmt.Printf("cpu seconds/herts: %d\n", mp.hertz)
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
		monproc = &process{path: path, pid: strconv.Itoa(PID)}
		monproc.getUptime()
		monproc.getCPUSeconds()
		monproc.getStat()

		// ** remove this when goroutines added ** //
		break
		// *************************************** //
	}
}

func main() {
	GetProcesses()
}
