Contains gRPC examples with
- unary
- client streaming
- server streaming
- bi-directional streaming
# Set Path
```
 export PATH="$PATH:$(go env GOPATH)/bin" 
```

# Generate Proto

```
 protoc --go_out=. --go-grpc_out=. ./proto/greet.proto
```