# go-connect-tcp

A simple TCP connection between a client and server in Golang.

Provide the `-p` flag to specify which port to run the server on. The default is `13102`.

Additionally provide the `-h` flag and you'll attempt to make a TCP connection to a given host.

For instance, the simplest way to create a local server + client:

```
$ go-connect-tcp # Creates the server at port 13102
$ go-connect-tcp -h localhost # Creates the client, connects to localhost:13102
```
