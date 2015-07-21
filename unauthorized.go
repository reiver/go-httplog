package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Unauthorized(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusUnauthorized
	httpStatusName :=  StatusNameUnauthorized

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
