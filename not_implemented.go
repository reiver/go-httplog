package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) NotImplemented(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusNotImplemented
	httpStatusName :=  StatusNameNotImplemented

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
