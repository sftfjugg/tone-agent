package entity

const APPName = "tone" + "agent"

// task status
const (
	TaskRunningStatus   = "running"
	TaskCompletedStatus = "completed"
)

// cron interval
const (
	HeartbeatInterval  = "*/10 * * * * *" // 每10s
	PullTaskInterval   = "*/3 * * * * *"  // 每3s
	SyncResultInterval = "*/5 * * * * *"  // 每5s
)

// dir
const (
	ResultDir = "/usr/share/" + APPName
	ConfigDir = "/etc/" + APPName
	LogDir    = "/var/log/" + APPName
)

// api
const (
	HeartbeatAPI  = "api/agent/heartbeat"
	PullTaskAPI   = "api/agent/pull_task"
	SyncResultAPI = "api/agent/result_sync"
)

// response
const (
	SuccessOk   = "ok"
	SuccessFail = "fail"
)
