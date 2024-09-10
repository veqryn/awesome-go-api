//go:build tools
// +build tools

package main

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
	_ "github.com/contiamo/openapi-generator-go/v2"
	_ "github.com/ogen-go/ogen/cmd/ogen"
)
