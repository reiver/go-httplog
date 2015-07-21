package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) UnsupportedMediaType(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusUnsupportedMediaType
	httpStatusName :=  StatusNameUnsupportedMediaType

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
