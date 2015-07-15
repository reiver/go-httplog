package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestedRangeNotSatisfiable(w http.ResponseWriter) {

	httpStatusCode := http.StatusRequestedRangeNotSatisfiable
	httpStatusName :=  StatusNameRequestedRangeNotSatisfiable

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
