package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) MethodNotAllowed(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusMethodNotAllowed
	httpStatusName :=  StatusNameMethodNotAllowed

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
