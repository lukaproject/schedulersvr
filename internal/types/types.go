// Code generated by goctl. DO NOT EDIT.
package types

type TaskContent struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	TaskType   string `json:"task_type"`
	Input      string `json:"input,optional"`
	Output     string `json:"output,optional"`
	CommitTime uint64 `json:"commit_time,optional"`
	BeginTime  uint64 `json:"begin_time,optional"`
	EndTime    uint64 `json:"end_time,optional"`
	Status     string `json:"status,optional"`
	WorkerId   string `json:"worker_id,optional"`
}

type WorkerContent struct {
	Id       string `json:"id,optional"`
	Name     string `json:"name"`
	TaskType string `json:"task_type"`
	Info     string `json:"info,optional"`
	Status   string `json:"status,optional"`
}

type AddTaskReq struct {
	SessionId string `json:"session_id"`
	Name      string `json:"name"`
	Input     string `json:"input"`
	TaskType  string `json:"task_type"`
}

type AddTaskResp struct {
	SessionId string `json:"session_id"`
	Name      string `json:"name"`
	TaskId    string `json:"task_id"`
}

type GetTaskReq struct {
	SessionId string `json:"session_id"`
	TaskId    string `json:"task_id"`
}

type GetTaskResp struct {
	SessionId string      `json:"session_id"`
	Task      TaskContent `json:"task"`
}

type FetchTaskReq struct {
	SessionId string `json:"session_id"`
	TaskType  string `json:"task_type"`
	WorkerId  string `json:"worker_id"`
}

type FetchTaskResp struct {
	SessionId string      `json:"session_id"`
	Task      TaskContent `json:"task"`
}

type UpdateTaskReq struct {
	SessionId string      `json:"session_id"`
	Task      TaskContent `json:"task"`
}

type AddWorkerReq struct {
	SessionId string        `json:"session_id"`
	Worker    WorkerContent `json:"worker"`
}

type GeneralResp struct {
	SessionId string `json:"session_id"`
	ErrorMsg  string `json:"error_msg,optional"`
}

type RemoveWorkerByNameReq struct {
	SessionId  string `json:"session_id"`
	WorkerName string `json:"worker_name"`
}
