go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

protoc --proto_path=./ --go_out=./ message.proto