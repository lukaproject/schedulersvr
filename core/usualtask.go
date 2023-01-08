package core

import "encoding/json"

type UsualTask struct {
	Id     string
	Type   string
	Input  string
	Output string
}

func (ut *UsualTask) GetId() string {
	return ut.Id
}

func (ut *UsualTask) GetType() string {
	return ut.Type
}

func (ut *UsualTask) Marshal() ([]byte, error) {
	return json.Marshal(ut)
}

func (ut *UsualTask) Unmarshal(b []byte) error {
	return json.Unmarshal(b, ut)
}

func (ut *UsualTask) MustMarshal() []byte {
	b, _ := ut.Marshal()
	return b
}

func (ut *UsualTask) MustUnmarshal(b []byte) {
	_ = ut.Unmarshal(b)
}

func (ut *UsualTask) SetOutput(output string) {
	ut.Output = output
}

func (ut *UsualTask) GetOutput() string {
	return ut.Output
}

func (ut *UsualTask) GetInput() string {
	return ut.Input
}
