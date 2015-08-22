# static-server

The replacement for Python http.server (formerly SimpleHTTPServer).

## Install

Get with `go`.

```sh
$ go get github.com/1000ch/static-server
```

Build this repository.

```sh
$ go build
```

Put the binary `static-server` to the directory in the PATH such like `/usr/local/bin`.

## Usage

Just type and execute `static-server` in the directory to serve.

```sh
$ static-server
```

You can pass a port argument. Default port is **8000**.

```sh
$ static-server -port 5000
```


## License

MIT: http://1000ch.mit-license.org