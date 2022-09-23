package entity

const (
	AgentAPISyncResult = "api/agent/result_sync"
	AgentAPIPullTask   = "api/agent/pull_task"
	AgentAPIHeartbeat  = "api/agent/heartbeat"

	AgentTSNSalt = "tone-agent"
)


type Config struct {
	TSN   string `json:"tsn"`
	Mode  string `json:"mode"`
	Proxy string `json:"proxy"`
	ToneProxy string `json:"tone_proxy"`
	TestLibProxy string `json:"testlib_proxy"`
}

type Server struct {
	IP  string `json:"ip"`
	TSN   string `json:"tsn"`
	Domain   string `json:"tone_proxy"`
}