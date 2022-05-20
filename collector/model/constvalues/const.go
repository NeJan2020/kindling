package constvalues

const (
	RequestCount        = "request_count"
	RequestTotalTime    = "request_total_time"
	ConnectTime         = "connect_time"
	RequestSentTime     = "request_sent_time"
	WaitingTtfbTime     = "waiting_ttfb_time"
	ContentDownloadTime = "content_download_time"

	RequestIo  = "request_io"
	ResponseIo = "response_io"

	SpanInfo = "KSpanInfo"
)

const (
	ProtocolHttp  = "http"
	ProtocolHttp2 = "http2"
	ProtocolGrpc  = "grpc"
	ProtocolDubbo = "dubbo"
	ProtocolDns   = "dns"
	ProtocolKafka = "kafka"
	ProtocolMysql = "mysql"
)

const (
	TcpstatEstablished = "Established"
	TcpstatSynSent     = "SynSent"
	TcpstatSynRecv     = "SynRecv"
	TcpstatFinWait1    = "FinWait1"
	TcpstatFinWait2    = "FinWait2"
	TcpstatTimeWait    = "TimeWait"
	TcpstatClose       = "Close"
	TcpstatCloseWait   = "CloseWait"
	TcpstatLastAck     = "LastAck"
	TcpstatListen      = "Listen"
	TcpstatClosing     = "Closing"
)
