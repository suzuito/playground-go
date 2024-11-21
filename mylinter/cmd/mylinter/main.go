package main

import (
	"github.com/suzuito/playground-go/mylinter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(mylinter.Analyzer) }
