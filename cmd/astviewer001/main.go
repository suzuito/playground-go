package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "a.out <package>")
		os.Exit(1)
	}
	dirPathPackage := os.Args[1]

	if err := parse(dirPathPackage); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func parse(
	dirPathPackage string,
) error {
	entries, err := os.ReadDir(dirPathPackage)
	if err != nil {
		return fmt.Errorf("os.ReadDir is failed: %w", err)
	}

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".go" {
			continue
		}

		filePath := filepath.Join(dirPathPackage, entry.Name())
		fset := token.NewFileSet()
		file, err := parser.ParseFile(
			fset,
			filePath,
			nil,
			0,
		)
		if err != nil {
			return fmt.Errorf("parser.ParseFile is failed: %w", err)
		}

		fmt.Println(filePath, file.Name, file.Pos(), file.End())

		ast.Print(fset, file)
	}

	return nil
}
