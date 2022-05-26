package transform

import (
	"github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/constant"
	v1 "github.com/Kindling-project/kindling/collector/consumer/exporter/flattenexporter/data/protogen/common/v1"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
)

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

func GenerateTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 4)
	//httpMethod
	httpMethod := labelsMap.GetStringValue(constlabels.HttpMethod)
	httpMethodValue := v1.KeyValue{Key: constant.HttpMethod, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: httpMethod}}}
	keySlice = append(keySlice, httpMethodValue)

	//HttpStatusCode
	httpStatusCode := labelsMap.GetIntValue(constlabels.HttpStatusCode)
	httpStatusCodeValue := v1.KeyValue{Key: constant.HttpStatusCode, Value: v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: httpStatusCode}}}
	keySlice = append(keySlice, httpStatusCodeValue)

	//httpUrl
	httpUrl := labelsMap.GetStringValue(constlabels.HttpUrl)
	httpUrlValue := v1.KeyValue{Key: constant.HttpUrl, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: httpUrl}}}
	keySlice = append(keySlice, httpUrlValue)

	//contentKey
	contentKey := labelsMap.GetStringValue(constlabels.ContentKey)
	contentKeyValue := v1.KeyValue{Key: constant.ContentKey, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: contentKey}}}
	keySlice = append(keySlice, contentKeyValue)
	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateKvlistValueSlice(key string, keySlice []v1.KeyValue, keyValueSlice *[]v1.KeyValue) *[]v1.KeyValue {
	keyValue := v1.KeyValue{Key: key, Value: v1.AnyValue{Value: &v1.AnyValue_KvlistValue{KvlistValue: &v1.KeyValueList{
		Values: keySlice,
	}}}}
	*keyValueSlice = append(*keyValueSlice, keyValue)
	return keyValueSlice
}

func GenerateKeyValueBytesSlice(key string, value []byte, keyValueSlice *[]v1.KeyValue) *[]v1.KeyValue {
	keyValue := v1.KeyValue{Key: key, Value: v1.AnyValue{Value: &v1.AnyValue_BytesValue{BytesValue: value}}}
	*keyValueSlice = append(*keyValueSlice, keyValue)
	return keyValueSlice
}

func GenerateArrayValueSlice(key string, value *v1.ArrayValue, keyValueSlice *[]v1.KeyValue) *[]v1.KeyValue {
	keyValue := v1.KeyValue{Key: key, Value: v1.AnyValue{Value: &v1.AnyValue_ArrayValue{ArrayValue: value}}}
	*keyValueSlice = append(*keyValueSlice, keyValue)
	return keyValueSlice
}

func generateArrayValue() *v1.ArrayValue {
	return &v1.ArrayValue{}
}
