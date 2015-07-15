package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) ExpectationFailed(w http.ResponseWriter) {

	httpStatusCode := http.StatusExpectationFailed
	httpStatusName :=  StatusNameExpectationFailed

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
