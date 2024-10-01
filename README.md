# zap-cli-logger

A Zap logger for easy use with CLI applications.

## Why?

Sometimes an HTTP service has multiple binaries bundled with it. For example,
there can be a package for the HTTP service itself, and some utility CLIs that
use an API client for that HTTP service to do things.

In this case, it's nice to have CLI printing that leverages the same dependency
as the HTTP server for logging to minimize the overall dependency footprint.

## Usage

See [the example](./example/) for how to set it up:

```
➜  zap-cli-logger git:(develop) ✗ go build -o build/example ./example
➜  zap-cli-logger git:(develop) ✗ ./build/example
you should always see this, and if -verbose is set you should see it with context
you should always see this with context	{"some": "context"}
main.main
	/Users/hugo.torres/projects/taxfyle/zap-cli-logger/example/main.go:22
runtime.main
	/opt/homebrew/Cellar/go/1.23.1/libexec/src/runtime/proc.go:272
you should always see this with context	{"some": "context"}
main.main
	/Users/hugo.torres/projects/taxfyle/zap-cli-logger/example/main.go:23
runtime.main
	/opt/homebrew/Cellar/go/1.23.1/libexec/src/runtime/proc.go:272
➜  zap-cli-logger git:(develop) ✗ ./build/example -verbose
you should only see this if -verbose is set, and you should see it with context	{"some": "context"}
you should always see this, and if -verbose is set you should see it with context	{"some": "context"}
you should always see this with context	{"some": "context"}
main.main
	/Users/hugo.torres/projects/taxfyle/zap-cli-logger/example/main.go:22
runtime.main
	/opt/homebrew/Cellar/go/1.23.1/libexec/src/runtime/proc.go:272
you should always see this with context	{"some": "context"}
main.main
	/Users/hugo.torres/projects/taxfyle/zap-cli-logger/example/main.go:23
runtime.main
	/opt/homebrew/Cellar/go/1.23.1/libexec/src/runtime/proc.go:272
```
