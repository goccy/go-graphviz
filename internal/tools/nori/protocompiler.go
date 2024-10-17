package nori

import (
	"context"
	"io"
	"os"

	"github.com/bufbuild/protocompile"
	"github.com/bufbuild/protocompile/linker"
	"github.com/bufbuild/protocompile/protoutil"
	"github.com/bufbuild/protocompile/reporter"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	_ "embed"
)

type errorReporter struct {
	errs []reporter.ErrorWithPos
}

func (r *errorReporter) Error(err reporter.ErrorWithPos) error {
	r.errs = append(r.errs, err)
	return nil
}

func (r *errorReporter) Warning(_ reporter.ErrorWithPos) {
}

func ProtoCompile(ctx context.Context, file string, importPaths ...string) ([]*descriptorpb.FileDescriptorProto, error) {
	var r errorReporter

	compiler := protocompile.Compiler{
		Resolver: protocompile.WithStandardImports(&protocompile.SourceResolver{
			ImportPaths: importPaths,
			Accessor: func(p string) (io.ReadCloser, error) {
				return os.Open(p)
			},
		}),
		SourceInfoMode: protocompile.SourceInfoStandard,
		Reporter:       &r,
	}
	files := []string{file}
	linkedFiles, err := compiler.Compile(ctx, files...)
	if err != nil {
		return nil, err
	}
	protoFiles := getProtoFiles(linkedFiles)
	return protoFiles, nil
}

func getProtoFiles(linkedFiles []linker.File) []*descriptorpb.FileDescriptorProto {
	var (
		protos         []*descriptorpb.FileDescriptorProto
		protoUniqueMap = make(map[string]struct{})
	)
	for _, linkedFile := range linkedFiles {
		for _, proto := range getFileDescriptors(linkedFile) {
			if _, exists := protoUniqueMap[proto.GetName()]; exists {
				continue
			}
			protos = append(protos, proto)
			protoUniqueMap[proto.GetName()] = struct{}{}
		}
	}
	return protos
}

func getFileDescriptors(file protoreflect.FileDescriptor) []*descriptorpb.FileDescriptorProto {
	var protoFiles []*descriptorpb.FileDescriptorProto
	fileImports := file.Imports()
	for i := 0; i < fileImports.Len(); i++ {
		protoFiles = append(protoFiles, getFileDescriptors(fileImports.Get(i))...)
	}
	protoFiles = append(protoFiles, protoutil.ProtoFromFileDescriptor(file))
	return protoFiles
}
