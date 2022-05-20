package tcpstatanalyzer

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"github.com/Kindling-project/kindling/collector/model/constnames"
	"github.com/Kindling-project/kindling/collector/model/constvalues"
	"github.com/hashicorp/go-multierror"
)

type TcpStat struct {
	// Count of TCP connections in state "Established"
	Established uint64
	// Count of TCP connections in state "Syn_Sent"
	SynSent uint64
	// Count of TCP connections in state "Syn_Recv"
	SynRecv uint64
	// Count of TCP connections in state "Fin_Wait1"
	FinWait1 uint64
	// Count of TCP connections in state "Fin_Wait2"
	FinWait2 uint64
	// Count of TCP connections in state "Time_Wait
	TimeWait uint64
	// Count of TCP connections in state "Close"
	Close uint64
	// Count of TCP connections in state "Close_Wait"
	CloseWait uint64
	// Count of TCP connections in state "Listen_Ack"
	LastAck uint64
	// Count of TCP connections in state "Listen"
	Listen uint64
	// Count of TCP connections in state "Closing"
	Closing uint64
}

func (a *TcpstatAnalyzer) Handle(pid int) error {
	var container = make(map[string]int)
	containerId, err := cgroupToContainerId(procRoot, pid)
	if containerId == "" || err != nil || container[containerId] > 0 {
		return err
	}
	container[containerId] = pid

	// handle tcp ipv4
	if err := a.handleTcp4(procRoot, pid, containerId); err != nil {
		return err
	}
	// handle tcp ipv6
	if err := a.handleTcp6(procRoot, pid, containerId); err != nil {
		return err
	}
	return nil
}

func (a *TcpstatAnalyzer) handleTcp4(procRoot string, pid int, containerId string) error {
	t, err := tcpStatsFromProc(procRoot, pid, "net/tcp")
	if err != nil {
		return fmt.Errorf("failure get tcp4 stats from pid[%d]: %v", pid, err)
	}

	gaugeGroup := newGaugeGroup(containerId, t)
	gaugeGroup.Labels.AddStringValue(constlabels.Mode, "tcp4")
	var retError error
	for _, nextConsumer := range a.consumers {
		err := nextConsumer.Consume(gaugeGroup)
		if err != nil {
			retError = multierror.Append(retError, err)
		}
	}
	return retError
}

func (a *TcpstatAnalyzer) handleTcp6(procRoot string, pid int, containerId string) error {
	t, err := tcpStatsFromProc(procRoot, pid, "net/tcp6")
	if err != nil {
		return fmt.Errorf("failure get tcp6 stats from pid[%d]: %v", pid, err)
	}

	gaugeGroup := newGaugeGroup(containerId, t)
	gaugeGroup.Labels.AddStringValue(constlabels.Mode, "tcp6")
	var retError error
	for _, nextConsumer := range a.consumers {
		err := nextConsumer.Consume(gaugeGroup)
		if err != nil {
			retError = multierror.Append(retError, err)
		}
	}
	return retError
}

func newGaugeGroup(containerId string, t TcpStat) *model.GaugeGroup {
	labels := model.NewAttributeMap()
	labels.AddStringValue(constlabels.ContainerId, containerId)
	var values = []*model.Gauge{
		{Name: constvalues.TcpstatEstablished, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.Established)}}},
		{Name: constvalues.TcpstatSynSent, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.SynSent)}}},
		{Name: constvalues.TcpstatSynRecv, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.SynRecv)}}},
		{Name: constvalues.TcpstatFinWait1, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.FinWait1)}}},
		{Name: constvalues.TcpstatFinWait2, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.FinWait2)}}},
		{Name: constvalues.TcpstatClose, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.Close)}}},
		{Name: constvalues.TcpstatClosing, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.Closing)}}},
		{Name: constvalues.TcpstatCloseWait, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.CloseWait)}}},
		{Name: constvalues.TcpstatLastAck, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.LastAck)}}},
		{Name: constvalues.TcpstatListen, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.Listen)}}},
		{Name: constvalues.TcpstatTimeWait, Data: &model.Gauge_Int{Int: &model.Int{Value: int64(t.TimeWait)}}},
	}
	gaugeGroup := model.NewGaugeGroup(constnames.TcpStatsGaugeGroup, labels, uint64(time.Now().UnixNano()), values...)
	return gaugeGroup
}

func cgroupToContainerId(procRoot string, pid int) (string, error) {
	cgroupFile := path.Join(procRoot, strconv.Itoa(pid), "cgroup")
	data, err := ioutil.ReadFile(cgroupFile)
	if err != nil {
		return "", fmt.Errorf("failure opening %s: %v", cgroupFile, err)
	}

	reader := strings.NewReader(string(data))
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ":")
		dir := fields[len(fields)-1]
		dirFields := strings.Split(dir, "/")
		containerId := dirFields[len(dirFields)-1]
		if isContainerId(containerId) {
			return containerId[0:12], nil
		}
	}
	return "", nil
}

func tcpStatsFromProc(procRoot string, pid int, file string) (TcpStat, error) {
	tcpStatsFile := path.Join(procRoot, strconv.Itoa(pid), file)

	tcpStats, err := scanTCPStats(tcpStatsFile)
	if err != nil {
		return tcpStats, fmt.Errorf("couldn't read tcp stats: %v", err)
	}

	return tcpStats, nil
}

func scanTCPStats(tcpStatsFile string) (TcpStat, error) {
	var stats TcpStat

	data, err := ioutil.ReadFile(tcpStatsFile)
	if err != nil {
		return stats, fmt.Errorf("failure opening %s: %v", tcpStatsFile, err)
	}

	tcpStateMap := map[string]uint64{
		"01": 0, // ESTABLISHED
		"02": 0, // SYN_SENT
		"03": 0, // SYN_RECV
		"04": 0, // FIN_WAIT1
		"05": 0, // FIN_WAIT2
		"06": 0, // TIME_WAIT
		"07": 0, // CLOSE
		"08": 0, // CLOSE_WAIT
		"09": 0, // LAST_ACK
		"0A": 0, // LISTEN
		"0B": 0, // CLOSING
	}

	reader := strings.NewReader(string(data))
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	// Discard header line
	if b := scanner.Scan(); !b {
		return stats, scanner.Err()
	}

	for scanner.Scan() {
		line := scanner.Text()

		state := strings.Fields(line)
		// TCP state is the 4th field.
		// Format: sl local_address rem_address st tx_queue rx_queue tr tm->when retrnsmt  uid timeout inode
		tcpState := state[3]
		_, ok := tcpStateMap[tcpState]
		if !ok {
			return stats, fmt.Errorf("invalid TCP stats line: %v", line)
		}
		tcpStateMap[tcpState]++
	}

	stats = TcpStat{
		Established: tcpStateMap["01"],
		SynSent:     tcpStateMap["02"],
		SynRecv:     tcpStateMap["03"],
		FinWait1:    tcpStateMap["04"],
		FinWait2:    tcpStateMap["05"],
		TimeWait:    tcpStateMap["06"],
		Close:       tcpStateMap["07"],
		CloseWait:   tcpStateMap["08"],
		LastAck:     tcpStateMap["09"],
		Listen:      tcpStateMap["0A"],
		Closing:     tcpStateMap["0B"],
	}

	return stats, nil
}

func isContainerId(str string) bool {
	reg := regexp.MustCompile(`^[a-z0-9]+$`)
	return reg.MatchString(str)
}
