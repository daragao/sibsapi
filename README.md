# SIBS API

Playing around with the [SIBS API](https://developer.sibsapimarket.com)

# WORK IN PROGRESS

The [`main.go`](main.go) is currently only used to print out the HTTP Request results (or parts of the structs marshalled).

The most interesting files at the moment arE:
 - [`client/types.go`](client/types.go) which holds the structs of the responses and requests payloads
 - [`client/client.go`](client/client.go) which has methods that do a request and return the body of the request in bytes
 
 _hopefully there will be more_
