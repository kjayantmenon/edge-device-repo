Asset Properties

- assetId
- name
- description
- protocol
- secondary_protocol
- primary_endpoint
- secondaryendpoint
- timeout

Protofile

#### Go Package

Package Name: _github.com/kjayantmenon/kjayantmenon/assetregistry_

#### Generate the protobuf files

1. Install the go packages

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
```

2. Generate the protobuf and grpc files

```bash

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
proto/AssetRegistry.proto

```
