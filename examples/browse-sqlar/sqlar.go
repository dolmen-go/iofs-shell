// Command browse-sqlar allows to browse an SQLite Archive file interactively on a terminal.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	shell "github.com/dolmen-go/iofs-shell"
	"github.com/dolmen-go/sqlar/sqlarfs"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.SetFlags(0)
	err := run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <file>.sqlar", filepath.Base(os.Args[0]))
	}

	db, err := sql.Open("sqlite3", "file:"+args[0]+"?mode=ro&immutable=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ar := sqlarfs.New(db, sqlarfs.PermOwner)
	return shell.Browse(ar, ".")
}
