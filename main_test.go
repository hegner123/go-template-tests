package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	tdb "github.com/hegner123/go-template-tests/sqlc"
)

func TestMain(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db, err := getDb()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	queries := tdb.New(db)
	rows, err := queries.ListAllPages(ctx)
	if err != nil {
		fmt.Println(err)
	}
	tr := NewTree(Scan, "ParseTree")
	for _, v := range rows {
		tr.Add(v)
		fmt.Println(v)
	}
	tr.Root.PrintTree(0)
}

func TestTemplate(t *testing.T) {
	b := "./templs"
	f, err := os.ReadFile(b + "/index.html")
	if err != nil {
		fmt.Println(err)
	}
    tr, err := os.ReadFile(b + "/menu/menu.html")
    if err != nil { 
        fmt.Println(err)
    }
    fmt.Printf("f: %s\n", f)
    fmt.Printf("f: %s\n", tr)
}



