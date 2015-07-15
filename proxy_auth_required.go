package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) ProxyAuthRequired(w http.ResponseWriter) {

	httpStatusCode := http.StatusProxyAuthRequired
	httpStatusName :=  StatusNameProxyAuthRequired

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
