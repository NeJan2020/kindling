package model

import (
	"bytes"
	"encoding/json"
)

const (
	IntMetricType MetricType = iota
	HistogramMetricType
	NoneMetricType
)

type MetricType int

type Metric struct {
	Name string
	//	Data can be assigned by:
	//	Int
	//	Histogram
	Data isMetricData
}

func (i *Metric) GetData() isMetricData {
	if i != nil {
		return i.Data
	}
	return nil
}

const MetricPrefix = "{\""
const MetricKey = "\":"
const MetricSuffix = "}"

func (i Metric) MarshalJSON() ([]byte, error) {
	if b, err := json.Marshal(i.Data); err != nil {
		return nil, err
	} else {
		var res bytes.Buffer
		res.WriteString(MetricPrefix)
		res.WriteString(i.Name)
		res.WriteString(MetricKey)
		res.Write(b)
		res.WriteString(MetricSuffix)
		return res.Bytes(), nil
	}
}

// TODO UnmarshalJSON
func (i *Metric) UnmarshalJSON(data []byte) error {
	panic("Need to implement!")
}

func (i *Metric) GetInt() *Int {
	if x, ok := i.GetData().(*Metric_Int); ok {
		return x.Int
	}
	return nil
}

func (i *Metric) GetHistogram() *Histogram {
	if x, ok := i.GetData().(*Metric_Histogram); ok {
		return x.Histogram
	}
	return nil
}

func (i *Metric) DataType() MetricType {
	switch i.GetData().(type) {
	case *Metric_Int:
		return IntMetricType
	case *Metric_Histogram:
		return HistogramMetricType
	default:
		return NoneMetricType
	}
}

func (i *Metric) Clear() {
	switch i.DataType() {
	case IntMetricType:
		i.GetInt().Value = 0
	case HistogramMetricType:
		histogram := i.GetHistogram()
		histogram.BucketCounts = nil
		histogram.Count = 0
		histogram.Sum = 0
		histogram.ExplicitBoundaries = nil
	}
}

type Int struct {
	Value int64
}

func NewIntMetric(name string, value int64) *Metric {
	return &Metric{Name: name, Data: &Metric_Int{Int: &Int{Value: value}}}
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

// TODO UnmarshalJSON
func (i Int) UnmarshalJSON(data []byte) error {
	panic("Need to implement!")
}

func NewMetric(name string, data isMetricData) *Metric {
	return &Metric{Name: name, Data: data}
}

type Histogram struct {
	Sum                int64
	Count              uint64
	ExplicitBoundaries []int64
	BucketCounts       []uint64
}

func NewHistogramMetric(name string, histogram *Histogram) *Metric {
	return &Metric{Name: name, Data: &Metric_Histogram{Histogram: histogram}}
}

type isMetricData interface {
	isMetricData()
}

func (*Metric_Int) isMetricData()       {}
func (*Metric_Histogram) isMetricData() {}

type Metric_Int struct {
	Int *Int
}

func (i Metric_Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Int)
}

// TODO UnmarshalJSON
func (i Metric_Int) UnmarshalJSON(data []byte) error {
	panic("Need to implement!")
}

type Metric_Histogram struct {
	Histogram *Histogram
}

func (h Metric_Histogram) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.Histogram)
}

// TODO UnmarshalJSON
func (h Metric_Histogram) UnmarshalJSON(data []byte) error {
	panic("Need to implement!")
}
