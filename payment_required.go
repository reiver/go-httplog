package httplog


import (
	"net/http"
)


func (httpLogger *internalHttpLogger) PaymentRequired(w http.ResponseWriter, cascade ...interface{}) {

	httpStatusCode := http.StatusPaymentRequired
	httpStatusName :=  StatusNamePaymentRequired

	httpLogger.jsonHttpResponse(w, httpStatusCode, httpStatusName, cascade...)
}
