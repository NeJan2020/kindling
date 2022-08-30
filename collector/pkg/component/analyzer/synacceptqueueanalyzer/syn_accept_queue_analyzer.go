package synacceptqueueanalyzer

import (
	"fmt"

	"os"

	"github.com/Kindling-project/kindling/collector/pkg/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/analyzer"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer"
	conntrackerpackge "github.com/Kindling-project/kindling/collector/pkg/metadata/conntracker"
	"github.com/Kindling-project/kindling/collector/pkg/model"
	"github.com/Kindling-project/kindling/collector/pkg/model/constlabels"
	"github.com/Kindling-project/kindling/collector/pkg/model/constnames"
	"github.com/hashicorp/go-multierror"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	SynAcceptQueueMetric analyzer.Type = "synacceptqueueanalyzer"
)

type SynAcceptQueueAnalyzer struct {
	consumers   []consumer.Consumer
	conntracker conntrackerpackge.Conntracker
	telemetry   *component.TelemetryTools
}

func NewSynAcceptQueueMetricAnalyzer(cfg interface{}, telemetry *component.TelemetryTools, nextConsumers []consumer.Consumer) analyzer.Analyzer {
	retAnalyzer := &SynAcceptQueueAnalyzer{
		consumers: nextConsumers,
		telemetry: telemetry,
	}
	conntracker, err := conntrackerpackge.NewConntracker(nil)
	if err != nil {
		telemetry.Logger.Warn("Conntracker cannot work as expected:", zap.Error(err))
	}
	retAnalyzer.conntracker = conntracker
	return retAnalyzer

}

func (a *SynAcceptQueueAnalyzer) Start() error {
	a.telemetry.Logger.Sugar().Infof("SynAcceptQueueAnalyzer Start...")
	return nil
}

func (a *SynAcceptQueueAnalyzer) ConsumableEvents() []string {
	return []string{
		constnames.TcpSynAcceptQueueEvent,
	}
}

// ConsumeEvent gets the event from the previous component
func (a *SynAcceptQueueAnalyzer) ConsumeEvent(event *model.KindlingEvent) error {
	var dataGroup *model.DataGroup
	var err error
	if event.Name == constnames.TcpSynAcceptQueueEvent {
		dataGroup, err = a.generateSynAcceptQueue(event)
	}

	if err != nil {
		if ce := a.telemetry.Logger.Check(zapcore.DebugLevel, "Event Skip, "); ce != nil {
			ce.Write(
				zap.Error(err),
			)
		}
		return nil
	}
	if dataGroup == nil {
		return nil
	}

	a.telemetry.Logger.Infof("[analyzer]SynAcceptQueueDataGroup: %s", dataGroup.String())

	var retError error
	for _, nextConsumer := range a.consumers {
		err := nextConsumer.Consume(dataGroup)
		if err != nil {
			retError = multierror.Append(retError, err)
		}
	}
	return retError
}

func (a *SynAcceptQueueAnalyzer) generateSynAcceptQueue(event *model.KindlingEvent) (*model.DataGroup, error) {
	// Only client-side has rtt metric
	labels, err := a.getSynAcceptQueueLabels(event)
	if err != nil {
		return nil, err
	}

	syn_len := event.GetIntUserAttribute("syn_len")
	syn_max := event.GetUintUserAttribute("syn_max")
	accept_len := event.GetUintUserAttribute("accept_len")
	accept_max := event.GetUintUserAttribute("accept_max")

	if ce := a.telemetry.Logger.Check(zapcore.DebugLevel, "Tcp Rtt: "); ce != nil {
		ce.Write(
			zap.Int64("syn_len", syn_len),
			zap.Uint64("syn_max", syn_max),
			zap.Uint64("accept_len", accept_len),
			zap.Uint64("accept_max", accept_max),
		)
	}

	syn_len_metric := model.NewIntMetric(constnames.TcpSynLenMetricName, int64(syn_len))
	syn_max_metric := model.NewIntMetric(constnames.TcpSynMaxMetricName, int64(syn_max))
	accept_len_metric := model.NewIntMetric(constnames.TcpAcceptLenMetricName, int64(accept_len))
	accept_max_metric := model.NewIntMetric(constnames.TcpAcceptMaxMetricName, int64(accept_max))

	var dataSlice []*model.Metric
	dataSlice = append(dataSlice, syn_len_metric)
	dataSlice = append(dataSlice, syn_max_metric)
	dataSlice = append(dataSlice, accept_len_metric)
	dataSlice = append(dataSlice, accept_max_metric)

	return model.NewDataGroup(constnames.TcpSynAcceptQueueMetricGroupName, labels, event.Timestamp, dataSlice...), nil
}

func (a *SynAcceptQueueAnalyzer) getSynAcceptQueueLabels(event *model.KindlingEvent) (*model.AttributeMap, error) {
	labels := model.NewAttributeMap()
	ctx := event.GetCtx()
	if ctx == nil {
		return labels, fmt.Errorf("ctx is nil for event %s", event.Name)
	}

	threadinfo := ctx.GetThreadInfo()
	if threadinfo == nil {
		return labels, fmt.Errorf("threadinfo is nil for event %s", event.Name)
	}

	tid := (int64)(threadinfo.GetTid())
	pid := (int64)(threadinfo.GetPid())

	containerId := threadinfo.GetContainerId()

	saddr := os.Getenv("MY_NODE_IP")
	sport := event.GetUintUserAttribute("sport")

	labels.AddIntValue(constlabels.Tid, tid)
	labels.AddIntValue(constlabels.Pid, pid)
	if containerId != "" {
		labels.AddStringValue(constlabels.ContainerId, containerId)
	}
	labels.AddStringValue(constlabels.NodeIp, saddr)
	labels.AddIntValue(constlabels.ListenPort, int64(sport))

	return labels, nil
}

// Shutdown cleans all the resources used by the analyzer
func (a *SynAcceptQueueAnalyzer) Shutdown() error {
	return nil
}

// Type returns the type of the analyzer
func (a *SynAcceptQueueAnalyzer) Type() analyzer.Type {
	return SynAcceptQueueMetric
}

func (a *SynAcceptQueueAnalyzer) SetSubEvents(params map[string]string) {

}
