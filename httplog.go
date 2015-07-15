package httplog


import (
	"io"
)


type internalHttpLogger struct {
	writer   io.Writer
	logs   []string
	dumpLogs bool
}


func New(writer io.Writer, dumpLogs bool) HttpLogger {
	logs := make([]string, 0, 16)

	httplogger := internalHttpLogger{
		writer:   writer,
		logs:     logs,
		dumpLogs: dumpLogs,
	}

	return &httplogger
}
