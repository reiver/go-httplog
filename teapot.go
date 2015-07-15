package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Teapot(w http.ResponseWriter) {

	httpStatusCode := http.StatusTeapot
	httpStatusName :=  StatusNameTeapot

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
