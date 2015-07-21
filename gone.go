package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Gone(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusGone
	httpStatusName :=  StatusNameGone

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
