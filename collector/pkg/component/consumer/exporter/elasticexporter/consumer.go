package elasticexporter

import (
	"github.com/Kindling-project/kindling/collector/pkg/model"
	"github.com/Kindling-project/kindling/collector/pkg/model/constnames"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func (e *ElasticExporter) Consume(dataGroup *model.DataGroup) error {
	for _, filter := range e.filters {
		if filter.IsDrop(dataGroup) {
			if ce := e.telemetry.Logger.Check(zapcore.DebugLevel, "data is filterd, skip: "); ce != nil {
				ce.Write(zap.String("Droped Data", dataGroup.String()))
			}
			return nil
		}
	}
	if dataGroup.Name == constnames.SingleNetRequestMetricGroup {
		return e.exportTrace(dataGroup)
	} else {
		return e.exportMetric(dataGroup)
	}
}
