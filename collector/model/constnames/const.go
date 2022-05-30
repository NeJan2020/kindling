package constnames

const (
	ReadEvent     = "read"
	WriteEvent    = "write"
	ReadvEvent    = "readv"
	WritevEvent   = "writev"
	SendToEvent   = "sendto"
	RecvFromEvent = "recvfrom"
	SendMsgEvent  = "sendmsg"
	RecvMsgEvent  = "recvmsg"

	TcpCloseEvent          = "tcp_close"
	TcpRcvEstablishedEvent = "tcp_rcv_established"
	TcpDropEvent           = "tcp_drop"
	TcpRetransmitSkbEvent  = "tcp_retransmit_skb"
	OtherEvent             = "other"

	GrpcUprobeEvent = "grpc_uprobe"
	// NetRequestMetricGroupName is used for dataGroup generated from networkAnalyzer.
	NetRequestMetricGroupName = "net_request_metric_group"
	// SingleNetRequestMetricGroup stands for the dataGroup with abnormal status.
	SingleNetRequestMetricGroup = "single_net_request_metric_group"
	// AggregatedNetRequestMetricGroup stands for the dataGroup after aggregation.
	AggregatedNetRequestMetricGroup = "aggregated_net_request_metric_group"

	TcpStatsGaugeGroup = "tcp_stats_gauge_group"

	TcpGaugeGroupName       = "tcp_metric_gauge_group"
	NodeGaugeGroupName      = "node_metric_gauge_group"
	PageFaultGaugeGroupName = "pgft_metric_gauge_group"
	ConnectGaugeGroupName   = "connect_metric_gauge_group"
	TcpMetricGroupName  = "tcp_metric_metric_group"
	NodeMetricGroupName = "node_metric_metric_group"
)
