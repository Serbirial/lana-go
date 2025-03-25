package error

func ErrorCheckPanic(e error) {
	if e != nil {
		panic(e)
	}
}
