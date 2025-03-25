package fs

import (
	"gobot/error"
	"os"
)

func ReadFileWhole(f string) string {
	data, err := os.ReadFile(f)
	error.ErrorCheckPanic(err)
	return string(data)
}
