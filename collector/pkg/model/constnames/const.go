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
	ConnectEvent  = "connect"

	PageFaultEvent   = "page_fault"
	SlowSyscallEvent = "slow_syscall"

	TcpSynAcceptQueueEvent = "inet_csk_accept"
	TcpCloseEvent          = "tcp_close"
	TcpRcvEstablishedEvent = "tcp_rcv_established"
	TcpDropEvent           = "tcp_drop"
	TcpRetransmitSkbEvent  = "tcp_retransmit_skb"
	TcpConnectEvent        = "tcp_connect"
	TcpSetStateEvent       = "tcp_set_state"
	OtherEvent             = "other"

	GrpcUprobeEvent = "grpc_uprobe"
	// NetRequestMetricGroupName is used for dataGroup generated from networkAnalyzer.
	NetRequestMetricGroupName = "net_request_metric_group"
	// SingleNetRequestMetricGroup stands for the dataGroup with abnormal status.
	SingleNetRequestMetricGroup = "single_net_request_metric_group"
	// AggregatedNetRequestMetricGroup stands for the dataGroup after aggregation.
	AggregatedNetRequestMetricGroup = "aggregated_net_request_metric_group"

	TcpMetricGroupName               = "tcp_metric_metric_group"
	NodeMetricGroupName              = "node_metric_metric_group"
	TcpConnectMetricGroupName        = "tcp_connect_metric_group"
	PgftMetricGroupName              = "pgft_metric_metric_group"
	TcpStatsMetricGroup              = "tcp_stats_metric_group"
	ErrorSlowSyscallGroupName        = "error_slow_syscall_trace_group"
	TcpSynAcceptQueueMetricGroupName = "inet_csk_accept_metric_group"
)
