package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) InternalServerError(w http.ResponseWriter) {

	httpStatusCode := http.StatusInternalServerError
	httpStatusName :=  StatusNameInternalServerError

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
