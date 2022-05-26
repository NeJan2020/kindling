package flattenexporter

import (
	"context"
	flattenTraces "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/trace/v1"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/common/v1"
	flattenMetrics "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/metrics/flatten"
	internalComponent "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/component"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/consumer"
	"go.uber.org/zap"
	"runtime"
	"sync"
	"time"
)

type Processor struct {
	logger           *zap.Logger
	exportCtx        context.Context
	timer            *time.Timer
	timeout          time.Duration
	sendBatchSize    int
	sendBatchMaxSize int
	newItem          chan interface{}
	batch            batch
	shutdownC        chan struct{}
	goroutines       sync.WaitGroup
}

type batch interface {
	// export the current batch
	export(ctx context.Context, sendBatchMaxSize int) error

	// itemCount returns the size of the current batch
	itemCount() int

	// add item to the current batch
	add(item interface{})
}

var _ TracesBatch = (*Processor)(nil)
var _ MetricsBatch = (*Processor)(nil)

func newBatchProcessor(cfg *Cfg, batch batch) (*Processor, error) {
	return &Processor{
		logger:           cfg.Telemetry.Logger,
		sendBatchSize:    int(cfg.Config.SendBatchSize),
		sendBatchMaxSize: int(cfg.Config.SendBatchMaxSize),
		timeout:          cfg.Config.BatchTimeout,
		newItem:          make(chan interface{}, runtime.NumCPU()),
		batch:            batch,
		shutdownC:        make(chan struct{}, 1),
		exportCtx:        context.Background(),
	}, nil
}

func (bp *Processor) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: true}
}

// Start is invoked during service startup.
func (bp *Processor) Start(context.Context, internalComponent.Host) error {
	bp.goroutines.Add(1)
	go bp.startProcessingCycle()
	return nil
}

// Shutdown is invoked during service shutdown.
func (bp *Processor) Shutdown(context.Context) error {
	close(bp.shutdownC)

	// Wait until all goroutines are done.
	bp.goroutines.Wait()
	return nil
}

func (bp *Processor) startProcessingCycle() {
	defer bp.goroutines.Done()
	bp.timer = time.NewTimer(bp.timeout)
	for {
		select {
		case <-bp.shutdownC:
		DONE:
			for {
				select {
				case item := <-bp.newItem:
					bp.processItem(item)
				default:
					break DONE
				}
			}
			// This is the close of the channel
			if bp.batch.itemCount() > 0 {
				// TODO: Set a timeout on sendTraces or
				// make it cancellable using the context that Shutdown gets as a parameter
				bp.sendItems()
			}
			return
		case item := <-bp.newItem:
			if item == nil {
				continue
			}
			bp.processItem(item)
		case <-bp.timer.C:
			if bp.batch.itemCount() > 0 {
				bp.sendItems()
			}
			bp.resetTimer()
		}
	}
}

func (bp *Processor) processItem(item interface{}) {
	bp.batch.add(item)
	sent := false

	for bp.batch.itemCount() >= bp.sendBatchSize {
		sent = true
		bp.sendItems()
	}

	if sent {
		bp.stopTimer()
		bp.resetTimer()
	}
}

func (bp *Processor) stopTimer() {
	if !bp.timer.Stop() {
		<-bp.timer.C
	}
}

func (bp *Processor) resetTimer() {
	bp.timer.Reset(bp.timeout)
}

func (bp *Processor) sendItems() {
	if err := bp.batch.export(bp.exportCtx, bp.sendBatchMaxSize); err != nil {
		bp.logger.Warn("Sender failed", zap.Error(err))
	}
}

func (bp *Processor) ConsumeTraces(_ context.Context, td flattenTraces.ExportTraceServiceRequest) error {
	bp.newItem <- td
	return nil
}

func (bp *Processor) ConsumeMetrics(_ context.Context, md flattenMetrics.FlattenMetrics) error {
	bp.newItem <- md
	return nil
}

type batchTraces struct {
	nextConsumer TracesBatch
	traceData    flattenTraces.ExportTraceServiceRequest
	count        int
}

func newBatchTraces(nextConsumer TracesBatch, initBatchTraces flattenTraces.ExportTraceServiceRequest) *batchTraces {
	return &batchTraces{nextConsumer: nextConsumer, traceData: initBatchTraces}
}

// add updates current batchTraces by adding new TraceData object
func (bt *batchTraces) add(item interface{}) {
	resourceSpan := item.(flattenTraces.ExportTraceServiceRequest)
	newResourceSpansCount := len(resourceSpan.ResourceSpans)
	if newResourceSpansCount == 0 {
		return
	}
	bt.count += newResourceSpansCount
	bt.traceData.ResourceSpans = append(bt.traceData.ResourceSpans, resourceSpan.ResourceSpans...)
}

func (bt *batchTraces) export(ctx context.Context, sendBatchMaxSize int) error {
	var req flattenTraces.ExportTraceServiceRequest
	if sendBatchMaxSize > 0 && bt.itemCount() > sendBatchMaxSize {
		req = SplitTraces(sendBatchMaxSize, &bt.traceData)
		bt.count -= sendBatchMaxSize
	} else {
		req = bt.traceData
		bt.traceData = InitTraceRequest()
		bt.count = 0
	}

	/*	for i := 0; i < len(req.ResourceSpans); i++ {
		for j := 0; j < len(req.ResourceSpans[i].InstrumentationLibrarySpans); j++ {
			fmt.Println(req.ResourceSpans[i].InstrumentationLibrarySpans[j].Spans[0].Events[0].Attributes[30].Value.GetValue())
		}
	}*/
	return bt.nextConsumer.ConsumeTraces(ctx, req)
}

func (bt *batchTraces) itemCount() int {
	return len(bt.traceData.ResourceSpans)
}

func newBatchMetrics(nextConsumer MetricsBatch, initBatchMetrics flattenMetrics.FlattenMetrics, service *v1.Service) *batchMetric {
	return &batchMetric{nextConsumer: nextConsumer, metricData: initBatchMetrics, service: service}
}

type batchMetric struct {
	nextConsumer MetricsBatch
	metricData   flattenMetrics.FlattenMetrics
	service      *v1.Service
	count        int
}

func (bm *batchMetric) itemCount() int {
	return len(bm.metricData.RequestMetricByte.Metrics)
}

func (bm *batchMetric) add(item interface{}) {
	metrics := item.(flattenMetrics.FlattenMetrics)
	newRequestMetricsCount := len(metrics.RequestMetricByte.Metrics)
	newMetricsCount := newRequestMetricsCount
	if (newMetricsCount) == 0 {
		return
	}
	bm.count += newMetricsCount
	bm.metricData.RequestMetricByte.Metrics = append(bm.metricData.RequestMetricByte.Metrics, metrics.RequestMetricByte.Metrics...)
}

func (bm *batchMetric) export(ctx context.Context, sendBatchMaxSize int) error {
	var req flattenMetrics.FlattenMetrics
	req = bm.metricData
	bm.metricData = InitMetricsRequest(bm.service)
	bm.count = 0
	return bm.nextConsumer.ConsumeMetrics(ctx, req)
}

// NewBatchTracesProcessor  creates a new batch processor that batches traces by size or with timeout
func NewBatchTracesProcessor(next TracesBatch, cfg *Cfg,
	initBatchTraces flattenTraces.ExportTraceServiceRequest) (*Processor, error) {
	return newBatchProcessor(cfg, newBatchTraces(next, initBatchTraces))
}

func NewBatchMetricsProcessor(next MetricsBatch, cfg *Cfg,
	initBatchMetrics flattenMetrics.FlattenMetrics) (*Processor, error) {
	return newBatchProcessor(cfg, newBatchMetrics(next, initBatchMetrics, cfg.Config.GetServiceInstance()))
}
