# grpc-practice

- [1000만 Traffic을 견디며 적용한 RPC 통신에 대해 학습하고 구현해보기](https://www.inflearn.com/course/%EB%8C%80%EC%9A%A9%EB%9F%89-tps-%EB%8C%80%EB%B9%84-rpc-%ED%86%B5%EC%8B%A0/dashboard)

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

