package context

// ------------------------------- Go Concurrency Patterns: Context -------------------------------
// Introduction
//
// In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional
// goroutines to access backends such as databases and RPC services. The set of goroutines working on a request
// typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the
// request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit
// quickly so the system can reclaim any resources they are using.
//
// At Google, we developed a context package that makes it easy to pass request-scoped values, cancellation signals, and
// deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly
// available as context. This article describes how to use the package and provides a complete working example.
//
// <editor-fold desc="Context">
//
// The core of the context package is the Context type:
//
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
//
//	type Context interface {
//		// Done returns a channel that is closed when this Context is canceled
//		// or times out.
//		Done() <-chan struct{}
//
//		// Err indicates why this context was canceled, after the Done channel
//		// is closed.
//		Err() error
//
//		// Deadline returns the time when this Context will be canceled, if any.
//		Deadline() (deadline time.Time, ok bool)
//
//		// Value returns the value associated with key or nil if none.
//		Value(key interface{}) interface{}
//	}
//
// (This description is condensed; the [godoc](https://pkg.go.dev/context) is authoritative.)
//
// The Done method returns a channel that acts as a cancellation signal to functions running on behalf of the Context:
// when the channel is closed, the functions should abandon their work and return. The Err method returns an error
// indicating why the Context was canceled. The [Pipelines and Cancellation](https://go.dev/blog/pipelines) article
// discusses the Done channel idiom in more detail.
//
// A Context does not have a Cancel method for the same reason the Done channel is receive-only: the function receiving
// a cancellation signal is usually not the one that sends the signal. In particular, when a parent operation starts
// goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Instead, the WithCancel
// function (described below) provides a way to cancel a new Context value.
//
// A Context is safe for simultaneous use by multiple goroutines. Code can pass a single Context to any number of
// goroutines and cancel that Context to signal all of them.
//
// The Deadline method allows functions to determine whether they should start work at all; if too little time is left,
// it may not be worthwhile. Code may also use a deadline to set timeouts for I/O operations.
//
// Value allows a Context to carry request-scoped data. That data must be safe for simultaneous use by multiple goroutines.
//
// </editor-fold>
//
// <editor-fold desc="Derived contexts">
//
// The context package provides functions to derive new Context values from existing ones. These values form a tree:
// when a Context is canceled, all Contexts derived from it are also canceled.
//
// Background is the root of any Context tree; it is never canceled:
//
// ```golang
// Background returns an empty Context. It is never canceled, has no deadline, and has no values. Background is
// typically used in main, init, and tests, and as the top-level Context for incoming requests.
// func Background() Context
// ```
//
// WithCancel and WithTimeout return derived Context values that can be canceled sooner than the parent Context.
// The Context associated with an incoming request is typically canceled when the request handler returns. WithCancel is
// also useful for canceling redundant requests when using multiple replicas. WithTimeout is useful for setting a
// deadline on requests to backend servers:
//
// // WithCancel returns a copy of parent whose Done channel is closed as soon as
// // parent.Done is closed or cancel is called.
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
//
// // A CancelFunc cancels a Context.
// type CancelFunc func()
//
// // WithTimeout returns a copy of parent whose Done channel is closed as soon as
// // parent.Done is closed, cancel is called, or timeout elapses. The new
// // Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// // any. If the timer is still running, the cancel function releases its
// // resources.
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
// WithValue provides a way to associate request-scoped values with a Context:
//
// // WithValue returns a copy of parent whose Value method returns val for key.
// func WithValue(parent Context, key interface{}, val interface{}) Context
// The best way to see how to use the context package is through a worked example.
// </editor-fold>
//
// <editor-fold desc="Adapting code for Contexts">
//
// Many server frameworks provide packages and types for carrying request-scoped values. We can define new
// implementations of the Context interface to bridge between code using existing frameworks and code that expects a
// Context parameter.

// For example, Gorilla’s github.com/gorilla/context package allows handlers to associate data with incoming requests by
// providing a mapping from HTTP requests to key-value pairs. In gorilla.go, we provide a Context implementation whose
// Value method returns the values associated with a specific HTTP request in the Gorilla package.
//
// Other packages have provided cancellation support similar to Context. For example,
// [Tomb](https://pkg.go.dev/gopkg.in/tomb.v2?utm_source=godoc) provides a Kill method that signals cancellation by
// closing a Dying channel. Tomb also provides methods to wait for those goroutines to exit, similar to sync.WaitGroup.
// In tomb.go, we provide a Context implementation that is canceled when either its parent Context is canceled or a
// provided Tomb is killed.
//
// </editor-fold>
//
// <editor-fold desc="Conclusion">

// At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the
// call path between incoming and outgoing requests. This allows Go code developed by many different teams to
// interoperate well. It provides simple control over timeouts and cancellation and ensures that critical values like
// security credentials transit Go programs properly.
//
// Server frameworks that want to build on Context should provide implementations of Context to bridge between their
// packages and those that expect a Context parameter. Their client libraries would then accept a Context from the
// calling code. By establishing a common interface for request-scoped data and cancellation, Context makes it easier
// for package developers to share code for creating scalable services.
// </editor-fold>
