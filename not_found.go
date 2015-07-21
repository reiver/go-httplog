package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) NotFound(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusNotFound
	httpStatusName :=  StatusNameNotFound

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
