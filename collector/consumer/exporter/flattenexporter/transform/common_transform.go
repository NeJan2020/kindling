package transform

import v1 "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/common/v1"

func GenerateKeyValueIntSlice(key string, value int64, keyValueSlice *[]v1.KeyValue) *[]v1.KeyValue {
	keyValue := v1.KeyValue{Key: key, Value: v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: value}}}
	*keyValueSlice = append(*keyValueSlice, keyValue)
	return keyValueSlice
}

func GenerateKeyValueStringSlice(key string, value string, keyValueSlice *[]v1.KeyValue) *[]v1.KeyValue {
	keyValue := v1.KeyValue{Key: key, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: value}}}
	*keyValueSlice = append(*keyValueSlice, keyValue)
	return keyValueSlice
}

func GenerateKeyValueBoolSlice(key string, value bool, keyValueSlice *[]v1.KeyValue) *[]v1.KeyValue {
	keyValue := v1.KeyValue{Key: key, Value: v1.AnyValue{Value: &v1.AnyValue_BoolValue{BoolValue: value}}}
	*keyValueSlice = append(*keyValueSlice, keyValue)
	return keyValueSlice
}

func GenerateStringKeyValueSlice(key string, value string, stringKeyValueSlice *[]v1.StringKeyValue) *[]v1.StringKeyValue {
	stringKeyValue := v1.StringKeyValue{Key: key, Value: value}
	*stringKeyValueSlice = append(*stringKeyValueSlice, stringKeyValue)
	return stringKeyValueSlice
}
