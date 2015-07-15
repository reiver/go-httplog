package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) BadGateway(w http.ResponseWriter) {

	httpStatusCode := http.StatusBadGateway
	httpStatusName :=  StatusNameBadGateway

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
