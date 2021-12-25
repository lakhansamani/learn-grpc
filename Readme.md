# Learning grpc with go

- Have golang installed
- Install protobuf : `brew install protobuf`
- Install go plugins

  - Install the protocol compiler plugins for Go using the following commands:

    ```sh
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    ```

  - Update your PATH so that the protoc compiler can find the plugins:

    ```sh
    	export PATH="$PATH:$(go env GOPATH)/bin"
    ```

- `make create`

- `cd server && go run server.go`

- `cd client && go run client.go`
