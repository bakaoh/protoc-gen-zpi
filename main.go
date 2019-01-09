package main

import (
	"flag"
	"fmt"

	"github.com/bakaoh/protoc-gen-zpi/extension"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/pseudomuto/protokit"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	// all the heavy lifting done for you!
	if err := protokit.RunPlugin(new(plugin)); err != nil {
		glog.Fatal(err)
	}
}

// plugin is an implementation of protokit.Plugin
type plugin struct{}

func (p *plugin) Generate(req *plugin_go.CodeGeneratorRequest) (*plugin_go.CodeGeneratorResponse, error) {
	glog.Error("Parsed code generator request")

	descriptors := protokit.ParseCodeGenRequest(req)

	resp := new(plugin_go.CodeGeneratorResponse)

	for _, d := range descriptors {
		fileName := *d.Name + ".zpi.go"
		content := "abc "

		for _, e := range d.EnumType {
			if i, err := proto.GetExtension(e.Options, extension.E_ZpiIn); err == nil {
				content += fmt.Sprintf("o %v\n", *i.(*bool))
			}

		}

		resp.File = append(resp.File, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(fileName),
			Content: proto.String(content),
		})
	}

	return resp, nil
}
