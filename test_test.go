package httplog


import (
	"testing"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)


// doTest does all the hard work for all the tests.
//
// The actual tests call this.
func doTest(t *testing.T, httpStatusCode int, httpStatusName string, fn func(HttpLogger, http.ResponseWriter)) {

	tests := []struct{
		Logs []string
	}{
		{
			Logs: []string{
				"apple",
			},
		},
		{
			Logs: []string{
				"apple",
				"banana",
			},
		},
		{
			Logs: []string{
				"apple",
				"banana",
				"cherry",
			},
		},
		{
			Logs: []string{
				"Hello world!",
				"The quick brown fox jumps over the lazy dog",
				"Lorem ipsum dolor. Sit amet neque enim dictum sed nulla etiam porttitor quisque ipsum sed orci ultrices vitae. Ipsum suscipit vitae. Non erat iure.",
				"Maecenas rhoncus donec. Varius ut tempus.",
				"Ut quam mi. Sit purus at vestibulum aenean eget. Iaculis magna mi a malesuada sit donec pulvinar id.",
				"Vestibulum class fermentum tellus odio in. Non duis dui vel sapien eu rutrum maecenas libero. Ornare sit vel.",
				"At porttitor ac vehicula in felis. Phasellus adipiscing donec. ",
			},
		},
	}


	for testNumber, test := range tests {

		var buffer bytes.Buffer

		dumpLogs := true
		httplogger := New(&buffer, dumpLogs)

		for _, log := range test.Logs {
			httplogger.Print(log)
		}

		mock := newMockResponseWriter()

		// Response specific.
		fn(httplogger, mock)

		// Response specific.
		if expected, actual := httpStatusCode, mock.StatusCode; expected != actual {
			t.Errorf("For test #%d, expected status code to be %d, but was actually %d.", testNumber, expected, actual)
			return
		}

		// Make sure the logs we write to it get dumped into the HTTP response header.
		//
		// And that they get dumped in the order we expected them to be dumped.
		for logNumber, log := range test.Logs {
			var  headersMap map[string][]string = mock.Headers

			key := fmt.Sprintf("X-Log.%03d", 1+logNumber)

			if _,ok := headersMap[key]; !ok {
				t.Errorf("For test #%d and log #%d, expected there be a header %q, but wasn't.", testNumber, logNumber, key)
				return
			}

			values := headersMap[key]

			if expected, actual := 1, len(values); expected != actual {
				t.Errorf("For test #%d and log #%d, expected there to be %d values for log %q, but was actually %d.", testNumber, logNumber, expected, key, actual)
				return
			}

			value := values[0]

			if expected, actual := log, value; expected != actual {
				t.Errorf("For test #%d and log #%d, expected log %q to be %q but was actually %q.", testNumber, logNumber, key, expected, actual)
				return
			}
		}

		// Make sure the logs sent the "normal" way are as expected.
		//
		// Note that because of the way we split the logs, we will actually get one extra line!
		// Because the last "\n" will create an extra empty line.
		// So we need len(test.Logs) == len(logLines)-1
		// Althuogh we correct this, dumping the last line.
		logLines := strings.Split(buffer.String(), "\n")
		logLines = logLines[0:len(logLines)-1]

		if actual, expected := len(logLines), len(test.Logs); expected != actual {
			t.Errorf("For test #%d, expected there to be %d logs outputted in the \"normal\" way, but actually was %d.\nExpected: %v\nActual %v", testNumber, expected, actual, test.Logs, logLines)
		}

		for logLineNumber, logLine := range logLines {
			if actual, expected := logLine, test.Logs[logLineNumber]; expected != actual {
				t.Errorf("For test #%d and log line #%d, expected log line to be %q, but actually was %q.", testNumber, logLineNumber, expected, actual)
			}
		}

		// Make logs HTTP response body is has some of the stuff we expect in there..
		//
		// For example:
		//	* it is sent as JSON,
		//	* it has a "status_code" field, and
		//	* it has a "status_name" field.
		//
		datum := struct{
			StatusCode int    `json:"status_code"`
			StatusName string `json:"status_name"`
		}{}

		responseBody := mock.Buffer.Bytes();

		if err := json.Unmarshal(responseBody, &datum); nil != err {
			t.Errorf("For test #%d, received an error when trying to unmarshal the response body as JSON: %v", testNumber, err)
		}

		// Response specific.
		if expected, actual := httpStatusCode, datum.StatusCode; expected != actual {
			t.Errorf("For test #%d, expected status code to be %d, but actually was %d.", testNumber, expected, actual)
		}

		// Response specific.
		if expected, actual := httpStatusName, datum.StatusName; expected != actual {
			t.Errorf("For test #%d, expected status name to be %q, but actually was %q.", testNumber, expected, actual)
		}
	}
}



func TestOK(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.OK(w)
	}

	doTest(t, http.StatusOK, StatusNameOK, fn)
}

func TestBadGateway(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.BadGateway(w)
	}

	doTest(t, http.StatusBadGateway, StatusNameBadGateway, fn)
}

func TestBadRequest(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.BadRequest(w)
	}

	doTest(t, http.StatusBadRequest, StatusNameBadRequest, fn)
}

func TestConflict(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.Conflict(w)
	}

	doTest(t, http.StatusConflict, StatusNameConflict, fn)
}

func TestExpectationFailed(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.ExpectationFailed(w)
	}

	doTest(t, http.StatusExpectationFailed, StatusNameExpectationFailed, fn)
}

func TestForbidden(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.Forbidden(w)
	}

	doTest(t, http.StatusForbidden, StatusNameForbidden, fn)
}

func TestGatewayTimeout(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.GatewayTimeout(w)
	}

	doTest(t, http.StatusGatewayTimeout, StatusNameGatewayTimeout, fn)
}

func TestGone(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.Gone(w)
	}

	doTest(t, http.StatusGone, StatusNameGone, fn)
}

func TestHTTPVersionNotSupported(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.HTTPVersionNotSupported(w)
	}

	doTest(t, http.StatusHTTPVersionNotSupported, StatusNameHTTPVersionNotSupported, fn)
}

func TestInternalServerError(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.InternalServerError(w)
	}

	doTest(t, http.StatusInternalServerError, StatusNameInternalServerError, fn)
}

func TestLengthRequired(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.LengthRequired(w)
	}

	doTest(t, http.StatusLengthRequired, StatusNameLengthRequired, fn)
}

func TestMethodNotAllowed(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.MethodNotAllowed(w)
	}

	doTest(t, http.StatusMethodNotAllowed, StatusNameMethodNotAllowed, fn)
}

func TestNotAcceptable(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.NotAcceptable(w)
	}

	doTest(t, http.StatusNotAcceptable, StatusNameNotAcceptable, fn)
}

func TestNotFound(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.NotFound(w)
	}

	doTest(t, http.StatusNotFound, StatusNameNotFound, fn)
}

func TestNotImplemented(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.NotImplemented(w)
	}

	doTest(t, http.StatusNotImplemented, StatusNameNotImplemented, fn)
}

func TestPaymentRequired(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.PaymentRequired(w)
	}

	doTest(t, http.StatusPaymentRequired, StatusNamePaymentRequired, fn)
}

func TestPreconditionFailed(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.PreconditionFailed(w)
	}

	doTest(t, http.StatusPreconditionFailed, StatusNamePreconditionFailed, fn)
}

func TestProxyAuthRequired(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.ProxyAuthRequired(w)
	}

	doTest(t, http.StatusProxyAuthRequired, StatusNameProxyAuthRequired, fn)
}

func TestRequestedRangeNotSatisfiable(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.RequestedRangeNotSatisfiable(w)
	}

	doTest(t, http.StatusRequestedRangeNotSatisfiable, StatusNameRequestedRangeNotSatisfiable, fn)
}

func TestRequestEntityTooLarge(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.RequestEntityTooLarge(w)
	}

	doTest(t, http.StatusRequestEntityTooLarge, StatusNameRequestEntityTooLarge, fn)
}

func TestRequestTimeout(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.RequestTimeout(w)
	}

	doTest(t, http.StatusRequestTimeout, StatusNameRequestTimeout, fn)
}

func TestRequestURITooLong(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.RequestURITooLong(w)
	}

	doTest(t, http.StatusRequestURITooLong, StatusNameRequestURITooLong, fn)
}

func TestServiceUnavailable(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.ServiceUnavailable(w)
	}

	doTest(t, http.StatusServiceUnavailable, StatusNameServiceUnavailable, fn)
}

func TestTeapot(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.Teapot(w)
	}

	doTest(t, http.StatusTeapot, StatusNameTeapot, fn)
}

func TestUnauthorized(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.Unauthorized(w)
	}

	doTest(t, http.StatusUnauthorized, StatusNameUnauthorized, fn)

}

func TestUnsupportedMediaType(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter) {
		httplogger.UnsupportedMediaType(w)
	}

	doTest(t, http.StatusUnsupportedMediaType, StatusNameUnsupportedMediaType, fn)
}
