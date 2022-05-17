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
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/config"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/component"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/consumer"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/consumer/consumererror"
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/internal/consumer/consumerhelper"

	"go.uber.org/zap"
)

type tracesRequest struct {
	baseRequest
	td     []byte
	pusher consumerhelper.ConsumeTracesFunc
}

func newTracesRequest(ctx context.Context, td []byte, pusher consumerhelper.ConsumeTracesFunc) request {
	return &tracesRequest{
		baseRequest: baseRequest{ctx: ctx},
		td:          td,
		pusher:      pusher,
	}
}

func (req *tracesRequest) onError(err error) request {
	var traceError consumererror.Traces
	if consumererror.AsTraces(err, &traceError) {
		return newTracesRequest(req.ctx, traceError.GetTraces(), req.pusher)
	}
	return req
}

func (req *tracesRequest) export(ctx context.Context) error {
	return req.pusher(ctx, req.td)
}

func (req *tracesRequest) count() int {
	return len(req.td)
}

type traceExporter struct {
	*baseExporter
	consumer.Traces
}

func NewTracesExporter(
	cfg config.Exporter,
	logger *zap.Logger,
	pusher consumerhelper.ConsumeTracesFunc,
	options ...Option,
) (component.TracesExporter, error) {

	if cfg == nil {
		return nil, errNilConfig
	}

	if logger == nil {
		return nil, errNilLogger
	}
	if pusher == nil {
		return nil, errNilPushTraceData
	}

	bs := fromOptions(options...)
	be := newBaseExporter(cfg, logger, bs)
	be.wrapConsumerSender(func(nextSender requestSender) requestSender {
		return &tracesExporterWithObservability{
			nextSender: nextSender,
		}
	})

	tc, err := consumerhelper.NewTraces(func(ctx context.Context, td []byte) error {
		return be.sender.send(newTracesRequest(ctx, td, pusher))
	}, bs.consumerOptions...)
	return &traceExporter{
		baseExporter: be,
		Traces:       tc,
	}, err
}

type tracesExporterWithObservability struct {
	//obsrep     *obsreport.Exporter
	nextSender requestSender
}

func (tewo *tracesExporterWithObservability) send(req request) error {
	//req.setContext(tewo.obsrep.StartTracesExportOp(req.context()))
	// Forward the data to the next consumer (this pusher is the next).
	err := tewo.nextSender.send(req)
	//tewo.obsrep.EndTracesExportOp(req.context(), req.count(), err)
	return err
}
