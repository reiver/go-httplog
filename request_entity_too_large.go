package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestEntityTooLarge(w http.ResponseWriter) {

	httpStatusCode := http.StatusRequestEntityTooLarge
	httpStatusName :=  StatusNameRequestEntityTooLarge

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
