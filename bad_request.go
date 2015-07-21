package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) BadRequest(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusBadRequest
	httpStatusName :=  StatusNameBadRequest

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
