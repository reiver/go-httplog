package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Gone(w http.ResponseWriter) {

	httpStatusCode := http.StatusGone
	httpStatusName :=  StatusNameGone

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
