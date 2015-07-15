package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) ServiceUnavailable(w http.ResponseWriter) {

	httpStatusCode := http.StatusServiceUnavailable
	httpStatusName :=  StatusNameServiceUnavailable

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
