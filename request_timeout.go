package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestTimeout(w http.ResponseWriter) {

	httpStatusCode := http.StatusRequestTimeout
	httpStatusName :=  StatusNameRequestTimeout

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
