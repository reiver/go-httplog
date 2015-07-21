package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) Conflict(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusConflict
	httpStatusName :=  StatusNameConflict

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
