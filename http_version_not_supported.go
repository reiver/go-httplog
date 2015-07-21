package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) HTTPVersionNotSupported(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusHTTPVersionNotSupported
	httpStatusName :=  StatusNameHTTPVersionNotSupported

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
