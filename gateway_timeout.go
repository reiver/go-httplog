package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) GatewayTimeout(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusGatewayTimeout
	httpStatusName :=  StatusNameGatewayTimeout

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
