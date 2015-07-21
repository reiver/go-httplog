package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Forbidden(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusForbidden
	httpStatusName :=  StatusNameForbidden

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
