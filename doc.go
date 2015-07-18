/*
Package httplog provides logging functionality that is more geared towards existing inside of an HTTP handler.

It provides 9 methods that you likely expect from a logger, which are exactly the same as what the ones in Go's
built-in "log" package. Namely: Fatal, Fatalf, Fatalln, Panic, Panicf, Panicln, Print, Printf and Println.

So, for example, you are able to do things such as:

	func (handler *internalHandler) ServeHTTP(w ResponseWriter, r *Request) {
		
		httplogger := httplog.New(io.Writer, true)
		
		httplog.Print("ServeHTTP() BEGIN")
		
		// ...
		
		httplog.Print("ServeHTTP() END")
	}

Which should be very familiar to anyone who has used Go's built in "log" package.

*/
package httplog
