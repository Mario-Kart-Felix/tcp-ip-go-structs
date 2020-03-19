# Primitive TCP/IP Server & Client implemented in Go

Based on <a href="https://appliedgo.net/networking/">this</a>.

Main file is `cmd/ad-hoc-tcp/networking.go`. It implements a `server` and a `client`.

1. Server, listens on the localhost port on all network interfaces (usually loopback and at least one real network card) using `net.Conn` wrapped with `bufio`'s `ReadWriter`

2. Client follows "ad-hoc" messaging protocol. It only sends strings like:

```
STRING\n
any string message here\n
```

and structs like:

```
GOB\n
any complexData struct
```

# How to run

To start the server

`$go run networking.go`

In a new terminal window run the client

`$go run networking.go --connect localhost`
