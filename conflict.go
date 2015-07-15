package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Conflict(w http.ResponseWriter) {

	httpStatusCode := http.StatusConflict
	httpStatusName :=  StatusNameConflict

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
