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

)
