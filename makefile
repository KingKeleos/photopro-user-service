.PHONY: gen_proto
gen_proto-gen: ## generate protobuf
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/user-service.proto