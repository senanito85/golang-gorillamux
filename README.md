# Simple Webapp based on Golang Gorilla/Mux framework 

Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

The name mux stands for "HTTP request multiplexer". Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions. The main features are:

It implements the http.Handler interface so it is compatible with the standard http.ServeMux.
Requests can be matched based on URL host, path, path prefix, schemes, header and query values, HTTP methods or using custom matchers.
URL hosts, paths and query values can have variables with an optional regular expression.
Registered URLs can be built, or "reversed", which helps maintaining references to resources.
Routes can be used as subrouters: nested routes are only tested if the parent route matches. This is useful to define groups of routes that share common conditions like a host, a path prefix or other repeated attributes. As a bonus, this optimizes request matching.

## Installation

With a correctly configured Go toolchain:

go get -u github.com/gorilla/mux

## Run This app with Docker

```bash
docker build -t senanito85/gorillamux:0.1.1 .
docker run -it --rm -p 8000:8000 senanito85/gorillamux:0.1.1  
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
