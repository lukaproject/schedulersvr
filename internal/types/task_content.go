package types

import (
	"encoding/json"

	"github.com/lukaproject/schedulersvr/gerrx"
)

func (tc *TaskContent) Serialize() []byte {
	b, err := json.Marshal(tc)
	gerrx.Must(err)
	return b
}

func (tc *TaskContent) Unserialize(b []byte) error {
	return json.Unmarshal(b, tc)
}
