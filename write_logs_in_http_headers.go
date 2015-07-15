package httplog


import (

	"net/http"
)


func (httplogger *internalHttpLogger) writeLogsInHttpHeaders(w http.ResponseWriter) {

	header := w.Header()

	for _, logMessage := range httplogger.logs {
		header.Set("X-Log", logMessage)
	}
}
