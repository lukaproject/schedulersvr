/*
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TaskTypeContent struct {
	Id string `json:"id"`
	Name string `json:"name"`
	MaxTaskInQueLimit int64 `json:"max_task_in_que_limit"`
	CreateTime int64 `json:"create_time"`
	ExtraInfo string `json:"extra_info,omitempty"`
}