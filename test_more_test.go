package httplog


import (
	"testing"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)



const (
	MoreTestExpectedOne   = "apple"
	MoreTestExpectedTwo   = "banana"
	MoreTestExpectedThree = "cherry"
	MoreTestExpectedFour  = "fig"
	MoreTestExpectedFive  = "kiwi"
	MoreTestExpectedSix   = "watermelon"
)


var (
	moreTestData = []interface{}{
		map[string]string {
			"one"   : MoreTestExpectedOne,
			"two"   : MoreTestExpectedTwo,
			"three" : MoreTestExpectedThree,
		},
		struct{
			Four string `json:"four" some:"junk"`
			Five string `what:"is" json:"five" this:"thing"`
			Six  string
		}{
			Four: MoreTestExpectedFour,
			Five: MoreTestExpectedFive,
			Six:  MoreTestExpectedSix,
		},
	}
)


// doMoreTest does all the hard work for all the tests.
//
// The actual tests call this.
func doMoreTest(t *testing.T, httpStatusCode int, httpStatusName string, fn func(HttpLogger, http.ResponseWriter, ...interface{}), theData ...interface{}) {

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
		fn(httplogger, mock, theData...)

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
		// Although we correct this, dumping the last line.
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

			One   string `json:"one"`
			Two   string `json:"two"`
			Three string `json:"three"`
			Four  string `json:"four"`
			Five  string `json:"five"`
			Six   string `json:"six"`
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


		// one.
		if expected, actual := MoreTestExpectedOne, datum.One; expected != actual {
			t.Errorf("For test #%d, expected \"one\" to be %q, but actually was %q.", testNumber, expected, actual)
		}

		// two.
		if expected, actual := MoreTestExpectedTwo, datum.Two; expected != actual {
			t.Errorf("For test #%d, expected \"two\" to be %q, but actually was %q.", testNumber, expected, actual)
		}

		// three.
		if expected, actual := MoreTestExpectedThree, datum.Three; expected != actual {
			t.Errorf("For test #%d, expected \"three\" to be %q, but actually was %q.", testNumber, expected, actual)
		}

		// four.
		if expected, actual := MoreTestExpectedFour, datum.Four; expected != actual {
			t.Errorf("For test #%d, expected \"four\" to be %q, but actually was %q.", testNumber, expected, actual)
		}

		// five.
		if expected, actual := MoreTestExpectedFive, datum.Five; expected != actual {
			t.Errorf("For test #%d, expected \"five\" to be %q, but actually was %q.", testNumber, expected, actual)
		}

		// Six.
		if expected, actual := MoreTestExpectedSix, datum.Six; expected != actual {
			t.Errorf("For test #%d, expected \"six\" to be %q, but actually was %q.", testNumber, expected, actual)
		}
	}
}


func TestMoreOK(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.OK(w, theData...)
	}

	doMoreTest(t, http.StatusOK, StatusNameOK, fn, moreTestData...)
}

func TestMoreBadGateway(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.BadGateway(w, theData...)
	}

	doMoreTest(t, http.StatusBadGateway, StatusNameBadGateway, fn, moreTestData...)
}

func TestMoreBadRequest(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.BadRequest(w, theData...)
	}

	doMoreTest(t, http.StatusBadRequest, StatusNameBadRequest, fn, moreTestData...)
}

func TestMoreConflict(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.Conflict(w, theData...)
	}

	doMoreTest(t, http.StatusConflict, StatusNameConflict, fn, moreTestData...)
}

func TestMoreExpectationFailed(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.ExpectationFailed(w, theData...)
	}

	doMoreTest(t, http.StatusExpectationFailed, StatusNameExpectationFailed, fn, moreTestData...)
}

func TestMoreForbidden(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.Forbidden(w, theData...)
	}

	doMoreTest(t, http.StatusForbidden, StatusNameForbidden, fn, moreTestData...)
}

func TestMoreGatewayTimeout(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.GatewayTimeout(w, theData...)
	}

	doMoreTest(t, http.StatusGatewayTimeout, StatusNameGatewayTimeout, fn, moreTestData...)
}

func TestMoreGone(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.Gone(w, theData...)
	}

	doMoreTest(t, http.StatusGone, StatusNameGone, fn, moreTestData...)
}

func TestMoreHTTPVersionNotSupported(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.HTTPVersionNotSupported(w, theData...)
	}

	doMoreTest(t, http.StatusHTTPVersionNotSupported, StatusNameHTTPVersionNotSupported, fn, moreTestData...)
}

func TestMoreInternalServerError(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.InternalServerError(w, theData...)
	}

	doMoreTest(t, http.StatusInternalServerError, StatusNameInternalServerError, fn, moreTestData...)
}

func TestMoreLengthRequired(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.LengthRequired(w, theData...)
	}

	doMoreTest(t, http.StatusLengthRequired, StatusNameLengthRequired, fn, moreTestData...)
}

func TestMoreMethodNotAllowed(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.MethodNotAllowed(w, theData...)
	}

	doMoreTest(t, http.StatusMethodNotAllowed, StatusNameMethodNotAllowed, fn, moreTestData...)
}

func TestMoreNotAcceptable(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.NotAcceptable(w, theData...)
	}

	doMoreTest(t, http.StatusNotAcceptable, StatusNameNotAcceptable, fn, moreTestData...)
}

func TestMoreNotFound(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.NotFound(w, theData...)
	}

	doMoreTest(t, http.StatusNotFound, StatusNameNotFound, fn, moreTestData...)
}

func TestMoreNotImplemented(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.NotImplemented(w, theData...)
	}

	doMoreTest(t, http.StatusNotImplemented, StatusNameNotImplemented, fn, moreTestData...)
}

func TestMorePaymentRequired(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.PaymentRequired(w, theData...)
	}

	doMoreTest(t, http.StatusPaymentRequired, StatusNamePaymentRequired, fn, moreTestData...)
}

func TestMorePreconditionFailed(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.PreconditionFailed(w, theData...)
	}

	doMoreTest(t, http.StatusPreconditionFailed, StatusNamePreconditionFailed, fn, moreTestData...)
}

func TestMoreProxyAuthRequired(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.ProxyAuthRequired(w, theData...)
	}

	doMoreTest(t, http.StatusProxyAuthRequired, StatusNameProxyAuthRequired, fn, moreTestData...)
}

func TestMoreRequestedRangeNotSatisfiable(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.RequestedRangeNotSatisfiable(w, theData...)
	}

	doMoreTest(t, http.StatusRequestedRangeNotSatisfiable, StatusNameRequestedRangeNotSatisfiable, fn, moreTestData...)
}

func TestMoreRequestEntityTooLarge(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.RequestEntityTooLarge(w, theData...)
	}

	doMoreTest(t, http.StatusRequestEntityTooLarge, StatusNameRequestEntityTooLarge, fn, moreTestData...)
}

func TestMoreRequestTimeout(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.RequestTimeout(w, theData...)
	}

	doMoreTest(t, http.StatusRequestTimeout, StatusNameRequestTimeout, fn, moreTestData...)
}

func TestMoreRequestURITooLong(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.RequestURITooLong(w, theData...)
	}

	doMoreTest(t, http.StatusRequestURITooLong, StatusNameRequestURITooLong, fn, moreTestData...)
}

func TestMoreServiceUnavailable(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.ServiceUnavailable(w, theData...)
	}

	doMoreTest(t, http.StatusServiceUnavailable, StatusNameServiceUnavailable, fn, moreTestData...)
}

func TestMoreTeapot(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.Teapot(w, theData...)
	}

	doMoreTest(t, http.StatusTeapot, StatusNameTeapot, fn, moreTestData...)
}

func TestMoreUnauthorized(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.Unauthorized(w, theData...)
	}

	doMoreTest(t, http.StatusUnauthorized, StatusNameUnauthorized, fn, moreTestData...)

}

func TestMoreUnsupportedMediaType(t *testing.T) {
	fn := func(httplogger HttpLogger, w http.ResponseWriter, theData ...interface{}) {
		httplogger.UnsupportedMediaType(w, theData...)
	}

	doMoreTest(t, http.StatusUnsupportedMediaType, StatusNameUnsupportedMediaType, fn, moreTestData...)
}
