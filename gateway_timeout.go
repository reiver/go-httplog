package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) GatewayTimeout(w http.ResponseWriter) {

	httpStatusCode := http.StatusGatewayTimeout
	httpStatusName :=  StatusNameGatewayTimeout

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
