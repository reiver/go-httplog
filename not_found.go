package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) NotFound(w http.ResponseWriter) {

	httpStatusCode := http.StatusNotFound
	httpStatusName :=  StatusNameNotFound

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
