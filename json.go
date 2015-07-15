package httplog


import (
	"encoding/json"
	"net/http"
)


func (httpLogger *internalHttpLogger) writeJsonHttpResponse(w http.ResponseWriter, httpResponseStatus int, datum interface{}) {

	// Specify the Content-Type of the HTTP response.
	w.Header().Set("Content-Type", "application/json")

	// Set the logs, which this HttpLogger received, as "X-Log" headers in
	// the HTTP response.
	httpLogger.setLogsInHttpHeaders(w)

	// Write the HTTP headers in the response, with the specified HTTP response code.
	w.WriteHeader(httpResponseStatus)

	// Write out the JSON response.
	jsonEncoder := json.NewEncoder(w)
	if nil == jsonEncoder {
		return
	}

	if err := jsonEncoder.Encode(datum); nil != err {
		return
	}
}



