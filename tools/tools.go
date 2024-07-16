//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "gotest.tools/gotestsum"
	_ "k8s.io/code-generator"
	_ "knative.dev/pkg/codegen/cmd/injection-gen"
	_ "mvdan.cc/gofumpt"
)
