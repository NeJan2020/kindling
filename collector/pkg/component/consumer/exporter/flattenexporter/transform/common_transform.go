package transform

import (
	"github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/constant"
	v1 "github.com/Kindling-project/kindling/collector/pkg/component/consumer/exporter/flattenexporter/data/protogen/common/v1"
	"github.com/Kindling-project/kindling/collector/pkg/model"
	"github.com/Kindling-project/kindling/collector/pkg/model/constlabels"
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

func GenerateHttpRequestTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 4)
	//httpMethod
	httpMethod := labelsMap.GetStringValue(constlabels.HttpMethod)
	httpMethodValue := v1.KeyValue{Key: constant.HttpMethod, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: httpMethod}}}
	keySlice = append(keySlice, httpMethodValue)

	//httpUrl
	httpUrl := labelsMap.GetStringValue(constlabels.HttpUrl)
	httpUrlValue := v1.KeyValue{Key: constant.HttpUrl, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: httpUrl}}}
	keySlice = append(keySlice, httpUrlValue)

	//contentKey
	contentKey := labelsMap.GetStringValue(constlabels.ContentKey)
	contentKeyValue := v1.KeyValue{Key: constant.ContentKey, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: contentKey}}}
	keySlice = append(keySlice, contentKeyValue)

	//request_payload
	requestPayload := labelsMap.GetStringValue(constlabels.HttpRequestPayload)
	payload := v1.KeyValue{Key: constant.RequestPayload, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: requestPayload}}}
	keySlice = append(keySlice, payload)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateHttpResponseTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 2)

	//HttpStatusCode
	httpStatusCode := labelsMap.GetIntValue(constlabels.HttpStatusCode)
	httpStatusCodeValue := v1.KeyValue{Key: constant.HttpStatusCode, Value: v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: httpStatusCode}}}
	keySlice = append(keySlice, httpStatusCodeValue)

	//response_payload
	responsePayload := labelsMap.GetStringValue(constlabels.HttpResponsePayload)
	payload := v1.KeyValue{Key: constant.ResponsePayload, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: responsePayload}}}
	keySlice = append(keySlice, payload)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateDubboRequestTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 1)
	//request_payload
	requestPayload := labelsMap.GetStringValue(constlabels.DubboRequestPayload)
	payload := v1.KeyValue{Key: constant.RequestPayload, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: requestPayload}}}
	keySlice = append(keySlice, payload)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateDubboResponseTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 2)

	//Dubbo Error_code
	errCode := labelsMap.GetIntValue(constlabels.DubboErrorCode)
	code := v1.KeyValue{Key: constant.DubboErrorCode, Value: v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: errCode}}}
	keySlice = append(keySlice, code)

	//response_payload
	responsePayload := labelsMap.GetStringValue(constlabels.DubboResponsePayload)
	payload := v1.KeyValue{Key: constant.ResponsePayload, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: responsePayload}}}
	keySlice = append(keySlice, payload)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateDNSRequestTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 1)
	//request_payload
	dnsDomain := labelsMap.GetStringValue(constlabels.DnsDomain)
	domain := v1.KeyValue{Key: constant.DNS_DOMAIN, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: dnsDomain}}}
	keySlice = append(keySlice, domain)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateDNSResponseTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 1)

	//DNS R_CODE
	rCode := labelsMap.GetIntValue(constlabels.DnsRcode)
	code := v1.KeyValue{Key: constant.DNS_R_CODE, Value: v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: rCode}}}
	keySlice = append(keySlice, code)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateKafkaRequestTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 1)
	//kafka Topic
	kafkaTopic := labelsMap.GetStringValue(constlabels.KafkaTopic)
	topic := v1.KeyValue{Key: constant.KafkaTopic, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: kafkaTopic}}}
	keySlice = append(keySlice, topic)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateKafkaResponseTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 1)

	//kafka ErrorCode
	errCode := labelsMap.GetIntValue(constlabels.KafkaErrorCode)
	code := v1.KeyValue{Key: constant.KafkaErrorCode, Value: v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: errCode}}}
	keySlice = append(keySlice, code)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateMysqlRequestTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 1)
	//mysql sql
	mysqlSql := labelsMap.GetStringValue(constlabels.Sql)
	sql := v1.KeyValue{Key: constant.MysqlSql, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: mysqlSql}}}
	keySlice = append(keySlice, sql)

	GenerateKvlistValueSlice(key, keySlice, keyValueSlice)
}

func GenerateMysqlResponseTraceApp(key string, labelsMap *model.AttributeMap, keyValueSlice *[]v1.KeyValue) {
	keySlice := make([]v1.KeyValue, 0, 2)

	//Mysql ErrorCode
	errCode := labelsMap.GetIntValue(constlabels.SqlErrCode)
	code := v1.KeyValue{Key: constant.MysqlErrorCode, Value: v1.AnyValue{Value: &v1.AnyValue_IntValue{IntValue: errCode}}}
	keySlice = append(keySlice, code)

	//Mysql ErrorMsg
	errMsg := labelsMap.GetStringValue(constlabels.SqlErrMsg)
	msg := v1.KeyValue{Key: constant.MysqlErrorMsg, Value: v1.AnyValue{Value: &v1.AnyValue_StringValue{StringValue: errMsg}}}
	keySlice = append(keySlice, msg)
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
