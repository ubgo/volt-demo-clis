// Command notes is one of two CLIs in the monorepo demo. Its releases are
// tagged notes/vX.Y.Z — the BINARY name, wherever this folder lives.
package main

import (
	"flag"
	"fmt"

	"github.com/ubgo/volt-demo-clis/internal/shared"
)

var version = "dev" // stamped via cmd/notes/.volt.yml

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("notes %s\n", version)
		return
	}
	fmt.Println(shared.Banner("notes"), "— your notes are safe")
}
