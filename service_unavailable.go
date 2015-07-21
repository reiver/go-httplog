package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) ServiceUnavailable(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusServiceUnavailable
	httpStatusName :=  StatusNameServiceUnavailable

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
