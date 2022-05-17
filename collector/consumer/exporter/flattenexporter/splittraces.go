package flattenexporter

import (
	flattenTraces "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/collector/trace/v1"
)

func SplitTraces(size int, src *flattenTraces.ExportTraceServiceRequest) flattenTraces.ExportTraceServiceRequest {
	if len(src.ResourceSpans) <= size {
		return *src
	}
	dest := flattenTraces.ExportTraceServiceRequest{
		ResourceSpans: src.ResourceSpans[0:size],
	}
	src.ResourceSpans = src.ResourceSpans[size:]
	return dest
}
