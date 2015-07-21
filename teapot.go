package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Teapot(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusTeapot
	httpStatusName :=  StatusNameTeapot

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
