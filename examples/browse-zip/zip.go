// Command browse-zip allows to browse a ZIP file interactively on a terminal.
package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"

	shell "github.com/dolmen-go/iofs-shell"
)

var prog = filepath.Base(os.Args[0])

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog)
	err := run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: %s <file>.zip", prog)
	}

	zr, err := zip.OpenReader(os.Args[1])
	if err != nil {
		return err
	}
	return shell.Browse(zr, ".")
}
