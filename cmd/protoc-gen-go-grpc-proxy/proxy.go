package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

const (
	contextPackage = protogen.GoImportPath("context")
	grpcPackage    = protogen.GoImportPath("google.golang.org/grpc")
)

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

// generateFileContent generates the gRPC service definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the grpc package it is being compiled against.")
	g.P("// Requires gRPC-Go v1.32.0 or later.")
	g.P("const _ = ", grpcPackage.Ident("SupportPackageIsVersion7")) // When changing, update version number above.
	g.P()
	for _, service := range file.Services {
		getServiceDesc(g, service)
	}
	g.P()
	for _, service := range file.Services {
		genService(g, service)
	}
}

func getServiceDesc(g *protogen.GeneratedFile, service *protogen.Service) {
	serverName := service.GoName + "Server"
	structName := service.GoName + "ClientProxy"
	g.P("var _ ", serverName, " = ", "(*", unexport(structName), ")(nil)")
	g.P()
	for _, m := range service.Methods {
		g.P("const GrpcOperation" + service.GoName + string(m.Desc.Name()) + " = " + fmt.Sprintf("/%s/%s", string(service.Desc.FullName()), string(m.Desc.Name())))
	}
}

func genService(g *protogen.GeneratedFile, service *protogen.Service) {
	clientName := service.GoName + "Client"
	serverName := service.GoName + "Server"
	structName := service.GoName + "ClientProxy"

	g.P("// ", unexport(structName), " is the proxy to turn ", service.GoName, " client to server interface.")
	g.P("//")

	// Client interface.
	g.P("type ", unexport(structName), " struct {")
	g.P("cc ", clientName)
	g.P("}")
	g.P()

	g.P("func New", structName, " (cc ", clientName, ") ", serverName, " {")
	g.P("return &", unexport(structName), "{cc}")
	g.P("}")
	g.P()

	for _, method := range service.Methods {

		g.P("func (c *", unexport(structName), ") ", serverSignature(g, method), "{")
		g.P("return c.cc.", method.GoName, "(ctx, in)")
		g.P("}")
	}

	g.P()
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	s := method.GoName + "(ctx " + g.QualifiedGoIdent(contextPackage.Ident("Context"))
	if !method.Desc.IsStreamingClient() {
		s += ", in *" + g.QualifiedGoIdent(method.Input.GoIdent)
	}
	s += ") ("
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		s += "*" + g.QualifiedGoIdent(method.Output.GoIdent)
	} else {
		s += method.Parent.GoName + "_" + method.GoName + "Client"
	}
	s += ", error)"
	return s
}

// generateFile generates a _grpc_proxy.pb.go file containing gRPC service definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}

	filename := file.GeneratedFilenamePrefix + "_grpc_proxy.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-grpc-proxy v", version)
	g.P("// - protoc             ", protocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }
