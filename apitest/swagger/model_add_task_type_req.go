/*
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AddTaskTypeReq struct {
	SessionId string `json:"session_id"`
	Name string `json:"name"`
	MaxTaskInQueLimit int64 `json:"max_task_in_que_limit"`
	ExtraInfo string `json:"extra_info,omitempty"`
}
