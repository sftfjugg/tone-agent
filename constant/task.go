package constant

type Task struct {
	Tid        string `json:"TID"`
	Script     string `json:"SCRIPT"`
	ScriptType string `json:"SCRIPT_TYPE"`
	Args       string `json:"ARGS"`
	Cwd        string `json:"CWD"`
	Timeout    int    `json:"TIMEOUT"`
	Env        string `json:"ENV"`
	Sync       bool   `json:"SYNC"`
}

type TaskResult struct {
	Success string `json:"SUCCESS"`
	Tasks   []Task `json:"TASKS"`
}

const (
	TaskRunningStatus   = "running"
	TaskCompletedStatus = "completed"
)
