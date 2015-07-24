package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) OK(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusOK
	httpStatusName :=  StatusNameOK

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
