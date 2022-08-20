package entity

type SuccessResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
}

type ErrorResponse struct {
	Code  string    `json:"code"`
	Msg   string `json:"msg"`
}

type AgentResponse struct {
	Tid        string `json:"TID"`
	Success    string `json:"SUCCESS"`
	TaskStatus string `json:"TASK_STATUS"`
	TaskResult string `json:"TASK_RESULT"`
	TaskPid    string `json:"TASK_PID"`
	ErrorCode  string `json:"ERROR_CODE"`
	ErrorMsg   string `json:"ERROR_MSG"`
	ExitCode   string `json:"EXIT_CODE"`
	FinishTime string `json:"FINISH_TIME"`
}

type TSNResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	TSN   string `json:"tsn"`
}

type IPResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	IP   string `json:"ip"`
}

type ConfigResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Config   Config `json:"config"`
}

type LogResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Log   string `json:"log"`
}
