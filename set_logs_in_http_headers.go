package httplog


import (
	"fmt"
	"net/http"
)


func (httplogger *internalHttpLogger) setLogsInHttpHeaders(w http.ResponseWriter) {

	if ! httplogger.dumpLogs {
		return
	}

	header := w.Header()

	for i, logMessage := range httplogger.logs {
		headerName := fmt.Sprintf("X-Log.%03d", 1+i)

		header.Set(headerName, logMessage)
	}
}
