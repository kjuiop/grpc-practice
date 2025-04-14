# grpc-practice

### Prerequisites

```go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```
export PATH="$PATH:$(go env GOPATH)/bin"
```


### Regenerate gRPC code

```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    gRPC/proto/auth.proto
```

- source_relative : proto type 이 있는 파일을 기준으로 gRPC code 생성
  - 해당 옵션이 없으면 파일 생성 위치가 Root Path 위치에서 생성됨

