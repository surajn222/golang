build:
	-brew install protobuf
	-go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	-protoc score.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --proto_path=.