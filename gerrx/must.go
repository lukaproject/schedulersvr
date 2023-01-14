package gerrx

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustPass[PassType any](ret PassType, err error) PassType {
	if err != nil {
		panic(err)
	}
	return ret
}
