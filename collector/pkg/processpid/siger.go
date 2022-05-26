package processpid

import (
	"fmt"
	"time"

	sigar "github.com/elastic/gosigar"
)

type Process struct {
	Pid          int          `json:"pid"`
	Ppid         int          `json:"ppid"`
	Name         string       `json:"name"`
	Username     string       `json:"username"`
	State        string       `json:"state"`
	CmdLine      string       `json:"cmdline"`
	Mem          *ProcMemStat `json:"mem"`
	Cpu          *ProcCpuTime `json:"cpu"`
	ctime        time.Time
	ListenPorts  []string `json:"listen_ports"`
	ProcessRoute string   `json:"process_route"`
}

type ProcMemStat struct {
	Size       uint64  `json:"size"`
	Rss        uint64  `json:"rss"`
	RssPercent float64 `json:"rss_p"`
	Share      uint64  `json:"share"`
}

type ProcCpuTime struct {
	User        uint64  `json:"user"`
	UserPercent float64 `json:"user_p"`
	System      uint64  `json:"system"`
	Total       uint64  `json:"total"`
	Start       string  `json:"start_time"`
}

func NewProcess(pid int) (*Process, error) {

	state := sigar.ProcState{}
	if err := state.Get(pid); err != nil {
		return nil, fmt.Errorf("error getting process state for pid=%d: %v", pid, err)
	}

	proc := Process{
		Pid:      pid,
		Ppid:     state.Ppid,
		Name:     state.Name,
		State:    getProcState(byte(state.State)),
		Username: state.Username,
		ctime:    time.Now(),
		//ListenPorts: config.PidPortsMap[strconv.Itoa(pid)],
	}

	return &proc, nil
}

func getProcState(b byte) string {

	switch b {
	case 'S':
		return "sleeping"
	case 'R':
		return "running"
	case 'D':
		return "idle"
	case 'T':
		return "stopped"
	case 'Z':
		return "zombie"
	}
	return "unknown"
}
