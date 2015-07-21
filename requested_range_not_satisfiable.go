package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) RequestedRangeNotSatisfiable(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusRequestedRangeNotSatisfiable
	httpStatusName :=  StatusNameRequestedRangeNotSatisfiable

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
