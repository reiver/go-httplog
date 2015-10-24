package httplog


import (
	"encoding/json"
	"net/http"
)


func (httpLogger *internalHttpLogger) jsonHttpResponse(w http.ResponseWriter, httpStatusCode int, httpStatusName string, cascade ...interface{}) {

	// Specify the Content-Type of the HTTP response.
	w.Header().Set("Content-Type", "application/json")

	// Set the logs, which this HttpLogger received, as "X-Log" headers in
	// the HTTP response.
	httpLogger.setLogsInHttpHeaders(w)

	// Write the HTTP headers in the response, with the specified HTTP response code.
	w.WriteHeader(httpStatusCode)

	// Collapse the cascade.
	data := collapse(cascade...)

	// Add the basic fields.
	//
	// These fields correspond to the HTTP response status code and name..
	data["status_code"] = httpStatusCode
	data["status_name"] = httpStatusName

	// Write out the JSON response.
	jsonEncoder := json.NewEncoder(w)
	if nil == jsonEncoder {
		return
	}

	if err := jsonEncoder.Encode(data); nil != err {
		return
	}
}
