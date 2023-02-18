package types

import "github.com/google/uuid"

func GenTaskTypeId() string {
	return uuid.NewString()
}

func GenTaskId() string {
	return uuid.NewString()
}
