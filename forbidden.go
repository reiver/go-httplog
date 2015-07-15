package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Forbidden(w http.ResponseWriter) {

	httpStatusCode := http.StatusForbidden
	httpStatusName :=  StatusNameForbidden

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
