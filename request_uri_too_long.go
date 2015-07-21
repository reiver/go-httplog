package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestURITooLong(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestURITooLong
	httpStatusName :=  StatusNameRequestURITooLong

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
