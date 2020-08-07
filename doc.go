// Package ctxsync allows threading a context through the blocking
// methods of the standard library's sync packge.
//
// All blocking methods from the standard library's sync package are implemented
// and based on channels instead.
//
// This may have a significant performance impact depending upon your use case.
//
// Benchmarks to follow.
package ctxsync
