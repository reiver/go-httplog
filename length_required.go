package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) LengthRequired(w http.ResponseWriter) {

	httpStatusCode := http.StatusLengthRequired
	httpStatusName :=  StatusNameLengthRequired

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
