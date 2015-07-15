package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) PaymentRequired(w http.ResponseWriter) {

	httpStatusCode := http.StatusPaymentRequired
	httpStatusName :=  StatusNamePaymentRequired

	datum := struct{
		StatusCode int    `json:"status_code"`
		StatusName string `json:"status_name"`
	}{
		StatusCode:    httpStatusCode,
		StatusName: httpStatusName,
	}

	httpLogger.writeJsonHttpResponse(w, httpStatusCode, datum)
}
