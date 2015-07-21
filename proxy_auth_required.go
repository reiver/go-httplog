package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) ProxyAuthRequired(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusProxyAuthRequired
	httpStatusName :=  StatusNameProxyAuthRequired

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
