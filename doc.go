/*
Package httplog provides logging functionality that is more geared towards existing inside of an HTTP handler.

It provides 9 methods that you likely expect from a logger, which are exactly the same the ones in Go's
built-in "log" package. Namely: Fatal, Fatalf, Fatalln, Panic, Panicf, Panicln, Print, Printf and Println.

So, for example, you are able to do things such as:

	func (handler *internalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
		
		httplogger := httplog.New(io.Writer, true)
		
		httplog.Print("ServeHTTP() BEGIN")
		
		// ...
		
		httplog.Print("ServeHTTP() END")
	}

Which should be very familiar to anyone who has used Go's built in "log" package.

However, httplog also provides 25 methods (that the Go's built-in "log" package does NOT have) that makes it
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

The usage of the InternalServerError method in the example (as well as any of the other 24 of the 25 methods
that make it easier to return a 4xx or 5xx HTTP response) will completely deal with send the HTTP response
over the 'http.ResponseWriter'.


*/
package httplog
