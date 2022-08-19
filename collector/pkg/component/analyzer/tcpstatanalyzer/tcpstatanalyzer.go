package tcpstatanalyzer

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Kindling-project/kindling/collector/pkg/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/analyzer"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer"
	"github.com/Kindling-project/kindling/collector/pkg/model"
	"go.uber.org/zap"
)

const (
	Tcpstat analyzer.Type = "tcpstatanalyzer"
)

type TcpstatAnalyzer struct {
	consumers []consumer.Consumer
	telemetry *component.TelemetryTools
	close     chan bool
	timestamp uint64
	container map[string]int
}

type Config struct {
}

const procRoot = "/proc"

func New(cfg interface{}, telemetry *component.TelemetryTools, consumers []consumer.Consumer) analyzer.Analyzer {
	return &TcpstatAnalyzer{
		consumers: consumers,
		telemetry: telemetry,
		container: make(map[string]int),
	}
}

// Start initializes the analyzer
func (a *TcpstatAnalyzer) Start() error {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for {
			select {
			case <-a.close:
				return
			case <-ticker.C:
				a.timestamp = uint64(time.Now().UnixNano())
				err := a.withAllProcs()
				if err != nil {
					a.telemetry.Logger.Error("Error collect tcp stats: ", zap.Error(err))
				}
				for k := range a.container {
					delete(a.container, k)
				}
			}
		}
	}()
	return nil
}

func (a *TcpstatAnalyzer) withAllProcs() error {
	files, err := ioutil.ReadDir(procRoot)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() || f.Name() == "." || f.Name() == ".." {
			continue
		}

		var pid int
		if pid, err = strconv.Atoi(f.Name()); err != nil {
			continue
		}

		if err = a.Handle(pid); err != nil {
			a.telemetry.Logger.Info("Unable to handle tcp stats: ", zap.Int("Pid", pid), zap.Error(err))
		}
	}
	return nil
}

// ConsumeEvent gets the event from the previous component
func (c *TcpstatAnalyzer) ConsumeEvent(event *model.KindlingEvent) error {
	return nil
}

// Shutdown cleans all the resources used by the analyzer
func (c *TcpstatAnalyzer) Shutdown() error {
	c.close <- true
	return nil
}

// Type returns the type of the analyzer
func (c *TcpstatAnalyzer) Type() analyzer.Type {
	return Tcpstat
}

func (c *TcpstatAnalyzer) ConsumableEvents() []string {
	return nil
}

func (c *TcpstatAnalyzer) SetSubEvents(params map[string]string) {

}
