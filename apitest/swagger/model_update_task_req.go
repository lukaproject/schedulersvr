/*
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type UpdateTaskReq struct {
	SessionId string `json:"session_id"`
	Task *TaskContent `json:"task"`
}
