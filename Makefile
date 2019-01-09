
genpb:
	protoc -I/usr/local/include -Iextension -I$$GOPATH/src \
		--go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:extension \
		extension/options.proto