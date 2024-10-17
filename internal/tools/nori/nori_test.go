package nori_test

import (
	"context"
	"path/filepath"
	"testing"

	"google.golang.org/protobuf/types/pluginpb"

	"github.com/goccy/nori"
)

func TestNori(t *testing.T) {
	bindProto := filepath.Join("internal", "wasm", "bind.proto")
	ctx := context.Background()
	files, err := nori.ProtoCompile(ctx, bindProto, filepath.Join(".."), filepath.Join("proto"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := nori.Generate(ctx, &pluginpb.CodeGeneratorRequest{
		ProtoFile: files,
	}); err != nil {
		t.Fatal(err)
	}
}
