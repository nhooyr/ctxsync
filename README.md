# ctxsync

This package implements the standard library's sync API but allows threading a context through
blocking methods.

Types without blocking methods are not duplicated here. Please use the standard library package
for those.

There is a performance hit as the types are implemented with channels instead of atomics but
I haven't benchmarked anything yet.

## Status

I haven't implemented it yet.

See [golang/go#40026](https://github.com/golang/go/issues/40026).
