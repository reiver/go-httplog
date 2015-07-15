package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) HTTPVersionNotSupported(w http.ResponseWriter) {

	httpStatusCode := http.StatusHTTPVersionNotSupported
	httpStatusName :=  StatusNameHTTPVersionNotSupported

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
