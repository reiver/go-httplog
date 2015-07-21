package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) PreconditionFailed(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusPreconditionFailed
	httpStatusName :=  StatusNamePreconditionFailed

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
