package entity

const (
	SuccessCode = 200
	SuccessMsg = "success"

	ParamsErrorCode = "a-1001"
	ParamsErrorMsg  = "The request parameters are incorrect. Please check if the parameters are correct"

	ExecCmdErrorCode = "a-1002"
	ExecCmdErrorMsg  = "script exec failed"

	PidNotExistErrorCode = "a-1003"
	PidNotExistErrorMsg  = "task pid no longer exists"

	SyncHeartbeatErrorCode = "b-1001"
	SyncHeartbeatErrorMsg  = "同步心跳失败，请检查配置"

	RequestErrorCode = "b-1002"
	RequestErrorMsg  = "请求失败，请检查URL配置"
)
