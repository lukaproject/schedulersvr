/*
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type FetchTaskReq struct {
	SessionId string `json:"session_id"`
	TaskType string `json:"task_type"`
	WorkerId string `json:"worker_id"`
}