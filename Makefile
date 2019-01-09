
extention:
	protoc -I/usr/local/include -Iextension -I$$GOPATH/src \
		--go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:extension \
		extension/options.proto

install:
	go install

example: extention install
	protoc -I/usr/local/include -Iexample -I$$GOPATH/src \
		--go_out=plugins=grpc,Mgithub.com/bakaoh/protoc-gen-zpi/extension/options.proto=github.com/bakaoh/protoc-gen-zpi/extension:example \
		--zpi_out=logtostderr=true:example \
		example/voucher.proto