package utils

import (
	"errors"
	"log/slog"
	"runtime"
)

func ErrorLog(message string, err error) {
	if err == nil {
		err = errors.New("some error")
	}
	
	// print file name and line of code where cause the error
	_, file, line, _ := runtime.Caller(1)
	slog.Error(
		message, "err", err.Error(),
		"file", file,
		"line", line)
}
