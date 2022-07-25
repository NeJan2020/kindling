package elasticexporter

import (
	"github.com/Kindling-project/kindling/collector/pkg/model"
	"github.com/Kindling-project/kindling/collector/pkg/model/constnames"
)

func (e *ElasticExporter) Consume(dataGroup *model.DataGroup) error {
	if dataGroup.Name == constnames.SingleNetRequestMetricGroup {
		return e.exportTrace(dataGroup)
	} else {
		return e.exportMetric(dataGroup)
	}
}
