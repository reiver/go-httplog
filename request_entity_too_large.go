package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestEntityTooLarge(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestEntityTooLarge
	httpStatusName :=  StatusNameRequestEntityTooLarge

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
