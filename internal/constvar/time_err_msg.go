package constvar

const (
	// 处理消息的超时时间 秒 time.Now().Unix() - c.Message().Unixtime
	TIME_OUT_MAX_seconds int64 = 2000
	// 服务故障错误提示
	ERR_MSG_Server string = "服务故障，可到交流群内反馈情况"
)
