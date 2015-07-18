/*
Package httplog provides logging functionality that is more geared towards existing inside of an HTTP handler.

It provides 9 methods that you likely expect from a logger, which are exactly the same the ones in Go's
built-in "log" package. Namely: Fatal, Fatalf, Fatalln, Panic, Panicf, Panicln, Print, Printf and Println.

So, for example, you are still able to do things such as:

	func (handler *internalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		
		httplogger := httplog.New(io.Writer, true)
		
		httplog.Print("ServeHTTP() BEGIN")
		
		// ...
		
		httplog.Print("ServeHTTP() END")
	}

Which should be very familiar to anyone who has used Go's built in "log" package.

However, httplog also provides 25 methods (that Go's built-in "log" package does NOT have) that makes it
easier to return a 4xx or 5xx HTTP response.

So, for example, you are able to do things such as:

	func (handler *internalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		
		httplogger := httplog.New(io.Writer, true)
		
		// ...
		
		if nil != err {
			httplogger.Printf("Received an error: %v", err)
			httplogger.InternalServerError(w)
			return
		}
		
		// ...
	}

The usage of the InternalServerError method in this example (as well as any of the other 24 of the 25 methods
that make it easier to return a 4xx or 5xx HTTP response) will completely deal with sending the HTTP response
over the 'http.ResponseWriter'.

The other important feature is that if the 'http logger' was created with its 'dumpLogs' parameters set to 'true'.
So, for example:

	httplogger := httplog.New(os.Stdout, true)

Then in the header of the HTTP response sent with any of the 25 methods (that make it easier to return a 4xx or 5xx
HTTP response), such as InternalServerError, NotFound, Unauthorized, NotImplemented, etc, you will get the logs dumped
using HTTP response headers that look like:

	x-Log.000: ServeHTTP() BEGIN
	x-Log.001: ServeHTTP() Received, subject = "This weekend?"
	x-Log.002: ServeHTTP() Received, body    = "Hey,\nLet's go to White Pine this weekend. Ok?"
	x-Log.003: ServeHTTP() Received, to      = ""
	x-Log.004: ServeHTTP() user_id = 5
	x-Log.005: ServeHTTP() Client Error: "to" is empty but is NOT allowed to be.
*/
package httplog
