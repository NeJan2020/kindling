package constant

const (
	Timestamp           = "timestamp"
	RequestCount        = "request_count"
	RequestTotalTime    = "request_total_time"
	ConnectTime         = "connect_time"
	ConnectFail         = "connect_fail"
	RequestSentTime     = "request_sent_time"
	WaitingTtfbTime     = "waiting_ttfb_time"
	ContentDownloadTime = "content_download_time"
	RequestIo           = "request_io"
	ResponseIo          = "response_io"
	RequestDurationTime = "request_duration_time"
	RequestAvgTime      = "request_avg_time"

	Protocol           = "protocol"
	APPProtocol        = "app_protocol"
	IsConnectFail      = "is_connect_fail"
	IsError            = "is_error"
	Status             = "status"
	SrcIp              = "src_ip"
	DstPort            = "dst_port"
	IsSlow             = "is_slow"
	IsServer           = "is_server"
	Pid                = "pid"
	RequestAPP         = "request_app"
	ContentKey         = "content_key"
	DstIp              = "dst_ip"
	SrcPort            = "src_port"
	ContainerId        = "container_id"
	ResponseAPP        = "response_app"
	SrcNode            = "src_node"
	SrcNamespace       = "src_namespace"
	SrcWorkloadKind    = "src_workload_kind"
	SrcWorkloadName    = "src_workload_name"
	SrcPod             = "src_pod"
	SrcPodIp           = "src_pod_ip"
	DstPodPort         = "dst_pod_port"
	DstContainerId     = "dst_containerid"
	DstContainer       = "dst_container"
	DstNode            = "dst_node"
	DstNamespace       = "dst_namespace"
	DstWorkloadName    = "dst_workload_name"
	DstWorkloadKind    = "dst_workload_kind"
	DstPod             = "dst_pod"
	DstPodIp           = "dst_pod_ip"
	SrcContainer       = "src_container"
	SrcContainerId     = "src_containerid"
	SrcMasterIP        = "src_masterip"
	DstMasterIP        = "dst_masterip"
	DstService         = "dst_service"
	DstServiceIp       = "dst_service_ip"
	SrcService         = "src_service"
	DstServicePort     = "dst_service_port"
	HttpHost           = "http_host"
	DNatIp             = "dnat_ip"
	DNatPort           = "dnat_port"
	Error              = "error"
	Slow               = "slow"
	StatusCode1xxTotal = "statuscode_1xx_total"
	StatusCode2xxTotal = "statuscode_2xx_total"
	StatusCode3xxTotal = "statuscode_3xx_total"
	StatusCode4xxTotal = "statuscode_4xx_total"
	StatusCode5xxTotal = "statuscode_5xx_total"

	Traces  = "traces"
	Metrics = "metrics"

	MetricTypeRequest    = 0
	MetricTypeConnect    = 1
	MetricTypeTcpStats   = 2
	MetricTypePageFault  = 3
	MetricTypeTcp        = 4
	MetricTypeSysCall    = 5
	MetricTypeTcpBacklog = 6
	MetricTypeTcpStatus  = 7

	ContainerName   = "container_name"
	DnsQueryTime    = "dns_query_time"
	RequestPayload  = "request_payload"
	ResponsePayload = "response_payload"
	HTTPS_TLS       = "https_TLS"
	DNS_R_CODE      = "dns_r_code"
	DNS_DOMAIN      = "dns_domain"
	MESSAGE_CAPTURE = "message_capture"

	HttpMethod       = "http_method"
	HttpUrl          = "http_url"
	HttpApmTraceType = "trace_type"
	HttpApmTraceId   = "trace_id"

	HttpStatusCode = "http_status_code"
	DubboErrorCode = "dubbo_error_code"

	KafkaErrorCode = "kafka_error_code"
	KafkaTopic     = "kafka_topic"

	MysqlSql       = "mysql_sql"
	MysqlErrorCode = "mysql_error_code"
	MysqlErrorMsg  = "mysql_error_message"
)