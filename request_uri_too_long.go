package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestURITooLong(w http.ResponseWriter) {

	httpStatusCode := http.StatusRequestURITooLong
	httpStatusName :=  StatusNameRequestURITooLong

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
