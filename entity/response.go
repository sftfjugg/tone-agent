package entity

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
	CODE int `json:"code"`
	MSG string `json:"msg"`
	TSN string `json:"tsn"`
}

type IPResponse struct {
	CODE int `json:"code"`
	MSG string `json:"msg"`
	IP string `json:"ip"`
}
