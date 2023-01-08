package core

type Task interface {
	GetId() string
	GetType() string
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	MustMarshal() []byte
	MustUnmarshal([]byte)
	GetInput() string
	SetOutput(string)
	GetOutput() string
}
