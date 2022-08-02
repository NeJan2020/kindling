// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporterhelper

import (
	"context"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/config"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/component"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/consumer"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/consumer/consumererror"
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/internal/consumer/consumerhelper"

	"go.uber.org/zap"
)

type metricsRequest struct {
	baseRequest
	md     []byte
	pusher consumerhelper.ConsumeMetricsFunc
}

func newMetricsRequest(ctx context.Context, md []byte, pusher consumerhelper.ConsumeMetricsFunc) request {
	return &metricsRequest{
		baseRequest: baseRequest{ctx: ctx},
		md:          md,
		pusher:      pusher,
	}
}

func (req *metricsRequest) onError(err error) request {
	var metricsError consumererror.Metrics
	if consumererror.AsMetrics(err, &metricsError) {
		return newMetricsRequest(req.ctx, metricsError.GetMetrics(), req.pusher)
	}
	return req
}

func (req *metricsRequest) export(ctx context.Context) error {
	return req.pusher(ctx, req.md)
}

func (req *metricsRequest) count() int {
	return len(req.md)
}

type metricsExporter struct {
	*baseExporter
	consumer.Metrics
}

// NewMetricsExporter creates an MetricsExporter that records observability metrics and wraps every request with a Span.
func NewMetricsExporter(
	cfg config.Exporter,
	logger *zap.Logger,
	pusher consumerhelper.ConsumeMetricsFunc,
	options ...Option,
) (component.MetricsExporter, error) {
	if cfg == nil {
		return nil, errNilConfig
	}

	if logger == nil {
		return nil, errNilLogger
	}

	if pusher == nil {
		return nil, errNilPushMetricsData
	}

	bs := fromOptions(options...)
	be := newBaseExporter(cfg, logger, bs)
	be.wrapConsumerSender(func(nextSender requestSender) requestSender {
		return &metricsSenderWithObservability{
			nextSender: nextSender,
		}
	})

	mc, err := consumerhelper.NewMetrics(func(ctx context.Context, md []byte) error {
		/*if bs.ResourceToTelemetrySettings.Enabled {
			md = convertResourceToLabels(md)
		}*/
		return be.sender.send(newMetricsRequest(ctx, md, pusher))
	}, bs.consumerOptions...)

	return &metricsExporter{
		baseExporter: be,
		Metrics:      mc,
	}, err
}

type metricsSenderWithObservability struct {
	nextSender requestSender
}

func (mewo *metricsSenderWithObservability) send(req request) error {
	//req.setContext(mewo.obsrep.StartMetricsExportOp(req.context()))
	err := mewo.nextSender.send(req)
	//mewo.obsrep.EndMetricsExportOp(req.context(), req.count(), err)
	return err
}
