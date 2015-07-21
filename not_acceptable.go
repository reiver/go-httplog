package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) NotAcceptable(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusNotAcceptable
	httpStatusName :=  StatusNameNotAcceptable

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
