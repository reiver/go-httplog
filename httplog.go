package httplog


import (
	"io"
)


type internalHttpLogger struct {
	writer io.Writer
	logs  []string
}


func New(writer io.Writer) HttpLogger {
	logs := make([]string, 0, 16)

	httplogger := internalHttpLogger{
		writer:  writer,
		logs:   logs,
	}

	return &httplogger
}
