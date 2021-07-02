package utils

func ErrorM(err error) {
	if err != nil {
		panic(err)
	}
}
