package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) LengthRequired(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusLengthRequired
	httpStatusName :=  StatusNameLengthRequired

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
