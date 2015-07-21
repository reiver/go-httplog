package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) InternalServerError(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusInternalServerError
	httpStatusName :=  StatusNameInternalServerError

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
