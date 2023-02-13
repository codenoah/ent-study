.PHONY: create
create:
	protoc -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/v1/*/*.proto

.PHONY: td-domain
td-domain:
	go test ./domain/... -v -coverprofile test-coverage.out
	go tool cover -func test-coverage.out

.PHONY: td-infra
td-infra:
	go test ./infra/... -v -coverprofile test-coverage.out
	go tool cover -func test-coverage.out

.PHONY: entity-create
entity-create:
	go run -mod=mod entgo.io/ent/cmd/ent init


.PHONY: init-protobuf
init-protobuf:
	go get -u google.golang.org/protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
