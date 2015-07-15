package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) UnsupportedMediaType(w http.ResponseWriter) {

	httpStatusCode := http.StatusUnsupportedMediaType
	httpStatusName :=  StatusNameUnsupportedMediaType

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
