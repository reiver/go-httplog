package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestTimeout(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestTimeout
	httpStatusName :=  StatusNameRequestTimeout

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
