package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"iter"
	"os"
)

func ReadLineFromFile(
	r io.Reader,
) iter.Seq[string] {
	return func(yield func(string) bool) {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

type ReturnedLine struct {
	Columns []string
	Err     error
}

func ReadLineFromSelectQuery(
	rows *sql.Rows,
) iter.Seq[ReturnedLine] {
	return func(yield func(ReturnedLine) bool) {
		for rows.Next() {
			cols, err := rows.Columns()
			if !yield(ReturnedLine{
				Columns: cols,
				Err:     err,
			}) {
				break
			}
		}
	}
}

func main() {
	for line := range ReadLineFromFile(os.Stdin) {
		fmt.Println(line)
	}
	rows := sql.Rows{}
	for row := range ReadLineFromSelectQuery(&rows) {
		if row.Err != nil {
			fmt.Println(row.Columns)
		}
	}
}
