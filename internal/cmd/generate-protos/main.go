// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run . -execute

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	gengogrpc "github.com/golang/protobuf/v2/cmd/protoc-gen-go-grpc/internal_gengogrpc"
	gengo "github.com/golang/protobuf/v2/cmd/protoc-gen-go/internal_gengo"
	"github.com/golang/protobuf/v2/protogen"
)

func init() {
	// When the environment variable RUN_AS_PROTOC_PLUGIN is set,
	// we skip running main and instead act as a protoc plugin.
	// This allows the binary to pass itself to protoc.
	if os.Getenv("RUN_AS_PROTOC_PLUGIN") != "" {
		protogen.Run(nil, func(gen *protogen.Plugin) error {
			for _, file := range gen.Files {
				if file.Generate {
					gengo.GenerateFile(gen, file)
					gengogrpc.GenerateFile(gen, file)
				}
			}
			return nil
		})
		os.Exit(0)
	}
}

var (
	run       bool
	protoRoot string
	repoRoot  string
)

func main() {
	flag.BoolVar(&run, "execute", false, "Write generated files to destination.")
	flag.StringVar(&protoRoot, "protoroot", os.Getenv("PROTOBUF_ROOT"), "The root of the protobuf source tree.")
	flag.Parse()
	if protoRoot == "" {
		panic("protobuf source root is not set")
	}

	// Determine repository root path.
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").CombinedOutput()
	check(err)
	repoRoot = strings.TrimSpace(string(out))

	generateLocalProtos()
	generateRemoteProtos()
}

func generateLocalProtos() {
	tmpDir, err := ioutil.TempDir(repoRoot, "tmp")
	check(err)
	defer os.RemoveAll(tmpDir)

	// Generate all local proto files (except version-locked files).
	dirs := []struct {
		path        string
		annotateFor map[string]bool
	}{
		{"cmd/protoc-gen-go/testdata", map[string]bool{"annotations/annotations.proto": true}},
		{"cmd/protoc-gen-go-grpc/testdata", nil},
		{"internal/testprotos", nil},
		{"encoding/testprotos", nil},
		{"reflect/protoregistry/testprotos", nil},
	}
	semVerRx := regexp.MustCompile(`v[0-9]+\.[0-9]+\.[0-9]+`)
	for _, d := range dirs {
		dstDir := filepath.Join(tmpDir, filepath.FromSlash(d.path))
		check(os.MkdirAll(dstDir, 0775))

		srcDir := filepath.Join(repoRoot, filepath.FromSlash(d.path))
		filepath.Walk(srcDir, func(srcPath string, _ os.FileInfo, _ error) error {
			if !strings.HasSuffix(srcPath, ".proto") || semVerRx.MatchString(srcPath) {
				return nil
			}
			relPath, err := filepath.Rel(srcDir, srcPath)
			check(err)

			// Emit a .meta file for certain files.
			var opts string
			if d.annotateFor[filepath.ToSlash(relPath)] {
				opts = ",annotate_code"
			}

			protoc("-I"+filepath.Join(protoRoot, "src"), "-I"+srcDir, "--go_out=paths=source_relative"+opts+":"+dstDir, relPath)
			return nil
		})
	}

	syncOutput(repoRoot, tmpDir)
}

func generateRemoteProtos() {
	tmpDir, err := ioutil.TempDir(repoRoot, "tmp")
	check(err)
	defer os.RemoveAll(tmpDir)

	// Generate all remote proto files.
	files := []struct{ prefix, path string }{
		{"", "conformance/conformance.proto"},
		{"src", "google/protobuf/any.proto"},
		{"src", "google/protobuf/api.proto"},
		{"src", "google/protobuf/descriptor.proto"},
		{"src", "google/protobuf/duration.proto"},
		{"src", "google/protobuf/empty.proto"},
		{"src", "google/protobuf/field_mask.proto"},
		{"src", "google/protobuf/source_context.proto"},
		{"src", "google/protobuf/struct.proto"},
		{"src", "google/protobuf/timestamp.proto"},
		{"src", "google/protobuf/type.proto"},
		{"src", "google/protobuf/wrappers.proto"},
		{"src", "google/protobuf/compiler/plugin.proto"},
	}
	for _, f := range files {
		protoc("-I"+filepath.Join(protoRoot, f.prefix), "--go_out="+tmpDir, f.path)
	}

	// Determine the module path.
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Path}}")
	cmd.Dir = repoRoot
	out, err := cmd.CombinedOutput()
	check(err)
	modulePath := strings.TrimSpace(string(out))

	syncOutput(repoRoot, filepath.Join(tmpDir, modulePath))
}

func protoc(args ...string) {
	cmd := exec.Command("protoc", "--plugin=protoc-gen-go="+os.Args[0])
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = append(os.Environ(), "RUN_AS_PROTOC_PLUGIN=1")
	cmd.Env = append(cmd.Env, "PROTOC_GEN_GO_ENABLE_REFLECT=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("executing: %v\n%s\n", strings.Join(cmd.Args, " "), out)
	}
	check(err)
}

func syncOutput(dstDir, srcDir string) {
	filepath.Walk(srcDir, func(srcPath string, _ os.FileInfo, _ error) error {
		if !strings.HasSuffix(srcPath, ".go") && !strings.HasSuffix(srcPath, ".meta") {
			return nil
		}
		relPath, err := filepath.Rel(srcDir, srcPath)
		check(err)
		dstPath := filepath.Join(dstDir, relPath)

		if run {
			fmt.Println("#", relPath)
			b, err := ioutil.ReadFile(srcPath)
			check(err)
			check(os.MkdirAll(filepath.Dir(dstPath), 0775))
			check(ioutil.WriteFile(dstPath, b, 0664))
		} else {
			cmd := exec.Command("diff", dstPath, srcPath, "-N", "-u")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
		return nil
	})
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
