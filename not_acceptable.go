package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) NotAcceptable(w http.ResponseWriter) {

	httpStatusCode := http.StatusNotAcceptable
	httpStatusName :=  StatusNameNotAcceptable

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
