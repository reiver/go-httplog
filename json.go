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

	// Deal with datum.
	datum := make(map[string]interface{})

	for _,x := range cascade {
		switch xx := x.(type) {
		case map[string]interface{}:
			for key, value := range xx {
				datum[key] = value
			}
		case map[string]string:
			for key, value := range xx {
				datum[key] = value
			}
		case string:
			datum["text"] = xx
		}
	}

	datum["status_code"] = httpStatusCode
	datum["status_name"] = httpStatusName

	// Write out the JSON response.
	jsonEncoder := json.NewEncoder(w)
	if nil == jsonEncoder {
		return
	}

	if err := jsonEncoder.Encode(datum); nil != err {
		return
	}
}
