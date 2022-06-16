# k8s-label-selector-shared-library

This repository demonstrates how to create a C shared library that exposes the label selector matching logic available in the Kubernetes source code.
Specifically, the shared library exposes the function `matches_label_selector`, which has the following function signature
```go
type matches_label_selector func(labels_json string, label_selector string) struct{matches bool, has_error bool, error_message string}
```

The function `matches_label_selector` invoke the same logic, available in the Go package [k8s.io/apimachinery/pkg/labels](https://pkg.go.dev/k8s.io/apimachinery/pkg/labels), that is executed when a request to the Kubernetes API contains a label selector.

A Ruby script is included that provides an example of consuming the shared library in Ruby.

## Usage

Building the Go main package in `./lib.go` generates a C header file <output>.h along with a C shared object <output>.so.

```shell
$ go build -o go_shared_lib.so -buildmode=c-shared lib.go
```

The Ruby file `main.rb` loads and consumes the shared library. It depends on the gem `ffi`, so make sure to fetch this dependency first e.g. `gem install ffi`.
Executing `./main.rb` invokes the shared
```shell
$ ruby ./main.rb
matches=true, has_error=false, error_message=
```
