.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest;
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest;


.PHONY: api
# 由 proto 生成接口层代码
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
	      user.proto