package graphviz_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/corona10/goimagehash"
	"github.com/goccy/go-graphviz"
)

var (
	testPaths = []string{
		filepath.Join("testdata", "directed"),
		filepath.Join("testdata", "undirected"),
	}
	imageHashJSON = filepath.Join("testdata", "imagehash.json")
)

const (
	imageThreshold = 15
)

func generateTestData() error {
	pathToHashDump := map[string]string{}
	for _, path := range testPaths {
		if err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			tmpfile, err := os.CreateTemp("", "graphviz")
			if err != nil {
				return err
			}
			defer os.Remove(tmpfile.Name())
			if err := exec.Command("dot", "-Tpng", fmt.Sprintf("-o%s", tmpfile.Name()), p).Run(); err != nil {
				return err
			}
			img, _, err := image.Decode(tmpfile)
			if err != nil {
				return err
			}
			hash, err := goimagehash.DifferenceHash(img)
			if err != nil {
				return err
			}
			var b bytes.Buffer
			if err := hash.Dump(&b); err != nil {
				return err
			}
			pathToHashDump[p] = base64.StdEncoding.EncodeToString(b.Bytes())
			return nil
		}); err != nil {
			return err
		}
	}
	content, err := json.Marshal(pathToHashDump)
	if err != nil {
		return err
	}
	if err := os.WriteFile(imageHashJSON, content, 0644); err != nil {
		return err
	}
	return nil
}

func TestGraphviz_Compatible(t *testing.T) {
	// generate testdata/imagehash.json
	//	if err := generateTestData(); err != nil {
	//		t.Fatal(err)
	//	}
	var pathToHashDump map[string]string
	file, err := os.ReadFile(imageHashJSON)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(file, &pathToHashDump); err != nil {
		t.Fatal(err)
	}
	for _, path := range testPaths {
		filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			file, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			graph, err := graphviz.ParseBytes(file)
			if err != nil {
				t.Fatal(err)
			}
			defer graph.Close()
			g := graphviz.New()
			defer g.Close()
			image, err := g.RenderImage(graph)
			if err != nil {
				t.Fatal(err)
			}
			hash, err := goimagehash.DifferenceHash(image)
			if err != nil {
				t.Fatal(err)
			}
			dump, err := base64.StdEncoding.DecodeString(pathToHashDump[path])
			if err != nil {
				t.Fatal(err)
			}
			targetHash, err := goimagehash.LoadImageHash(bytes.NewBuffer(dump))
			if err != nil {
				t.Fatal(err)
			}
			distance, err := hash.Distance(targetHash)
			if err != nil {
				t.Fatal(err)
			}
			if distance > imageThreshold {
				t.Fatalf("doesn't compatible image with dot. %s distance = %d", path, distance)
			}
			return nil
		})
	}
}
