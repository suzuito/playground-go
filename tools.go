//go:build tools

package tools

import (
	_ "github.com/gostaticanalysis/skeleton/v2"
	_ "golang.org/x/tools/cmd/callgraph"
	_ "golang.org/x/tools/cmd/goimports"
)
