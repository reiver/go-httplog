package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) ExpectationFailed(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusExpectationFailed
	httpStatusName :=  StatusNameExpectationFailed

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
