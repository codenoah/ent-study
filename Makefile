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


.PHONY: protobuf-download
protobuf-download:
	go get -u google.golang.org/protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


.PHONY: ent-gen
ent-gen:
	go generate ./ent