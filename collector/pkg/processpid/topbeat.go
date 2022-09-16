package processpid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Kindling-project/kindling/collector/pkg/component"
	"github.com/Kindling-project/kindling/collector/pkg/processpid/procs"
	"go.uber.org/zap"
)

var AllPidPortMap map[uint16]PortPidMapping

type PortPidMapping struct {
	Port uint16
	Pid  int
	Proc *Process
}

type PortPidMessage struct {
	MasterIp   string `json:"masterIp"`
	LogType    string `json:"logType"`
	Source     string `json:"source"`
	LogMessage MapStr `json:"logMessage"`
}

var client *http.Client

const processPidPortPostPath = "processpidport"

func InitSendPidPortBytime(interval time.Duration, address string, masterIp string, hostIp string, tools *component.TelemetryTools) {
	endpoint := fmt.Sprintf("%s/%s", address, processPidPortPostPath)
	go SendPidPortBytime(interval, endpoint, masterIp, hostIp, tools)
	tools.Logger.Info("ProcessPidPort Fetch Start!", zap.Float64("interval(s)", interval.Seconds()), zap.String("endpoint", endpoint), zap.String("masterIp ", masterIp), zap.String("hostIp", hostIp))
}
func SendPidPortBytime(interval time.Duration, endpoint string, masterIp string, hostIp string, tools *component.TelemetryTools) {
	if interval == 0 {
		interval = 15 * time.Second
	}

	host := os.Getenv("HCMINE_HOST_ROOT")
	UpdateAllPidMap(host)
	e := MapStr{
		"@timestamp": Time(time.Now()),
		"type":       "processPidPort",
		"dataMap":    AllPidPortMap,
		"HostIp":     hostIp,
	}

	PublishAllPidPort(e, endpoint, masterIp)
	procs.Pids = procs.Pids[0:0]
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			UpdateAllPidMap(host)
			e := MapStr{
				"@timestamp": Time(time.Now()),
				"type":       "processPidPort",
				"dataMap":    AllPidPortMap,
				"HostIp":     hostIp,
			}
			PublishAllPidPort(e, endpoint, masterIp)
			procs.Pids = procs.Pids[0:0]
		}
	}
}

func PublishAllPidPort(e MapStr, endpoint string, masterIp string) {
	msg := PortPidMessage{
		MasterIp:   masterIp,
		LogType:    "processPidPort",
		Source:     "hcmine-pipeline",
		LogMessage: e,
	}
	msgBytes, _ := json.Marshal(msg)
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Printf("error occurred when creating post request to %s:%+v\n", endpoint, err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	if client == nil {
		client = createHTTPClient()
	}
	post, err := client.Do(request)
	if post != nil && post.Body != nil {
		defer post.Body.Close()
	}
	if err != nil {
		log.Printf("error occurred when posting to %s:%+v\n", endpoint, err)
		return
	}
}

func UpdateAllPidMap(host string) {
	AllPidPortMap = make(map[uint16]PortPidMapping)
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(fmt.Sprintf("%s/proc", host))
	if err != nil {
		log.Println(err)
	}
	for i := range fileInfoList {
		if procs.IsNum(fileInfoList[i].Name()) {
			procs.Pids = append(procs.Pids, fileInfoList[i].Name())
		}

	}

	file, err := os.Open(fmt.Sprintf("%s/proc/net/tcp", host))
	if err != nil {
		log.Printf("Open: %s\n", err)
		return
	}
	defer file.Close()
	fileTcp6, err := os.Open(fmt.Sprintf("%s/proc/net/tcp6", host))
	if err != nil {
		log.Printf("Open: %s\n", err)
		return
	}
	defer fileTcp6.Close()
	socks, err := procs.Parse_Proc_Net_Tcp(file, false)
	socksTcp6, err := procs.Parse_Proc_Net_Tcp(fileTcp6, true)

	if err != nil {
		log.Printf("Parse_Proc_Net_Tcp: %s\n", err)
		return
	}
	socks_map := map[int64]*procs.SocketInfo{}
	for _, s := range socks {
		socks_map[s.Inode] = s
	}
	for _, s := range socksTcp6 {
		socks_map[s.Inode] = s
	}

	for _, pid := range procs.Pids {
		pidint, _ := strconv.Atoi(pid)
		inodes, err := procs.FindSocketsOfPid(host, pidint)
		if err != nil {
			log.Printf("FindSocketsOfPid: %s", err)
			continue
		}

		for _, inode := range inodes {
			sockInfo, exists := socks_map[inode]
			if exists {
				process, _ := NewProcess(pidint)
				entry := PortPidMapping{Port: sockInfo.Src_port, Pid: pidint, Proc: process}
				AllPidPortMap[sockInfo.Src_port] = entry
			}
		}
	}
}

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 90
)

func createHTTPClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
	}
	return client
}
