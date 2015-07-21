package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) BadGateway(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusBadGateway
	httpStatusName :=  StatusNameBadGateway

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
